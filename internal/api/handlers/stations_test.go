package handlers

import (
	"bysykkel/internal/api/handlers/mocks"
	"bysykkel/internal/api/restapi/operations"
	"bysykkel/internal/clients"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestStations_Handle(t *testing.T) {
	ctx := context.Background()
	t.Run("Fails to get station info", func(t *testing.T) {
		stations, byMock := setup(t)
		byMock.EXPECT().GetStationInfo(ctx).Return(nil, fmt.Errorf("err")).Times(1)

		rawReq := httptest.NewRequest("GET", "/", nil).WithContext(ctx)
		params := operations.GetStationsInfoParams{
			HTTPRequest: rawReq,
		}
		resp := stations.Handle(params)
		assert.IsType(t, &operations.GetStationsInfoInternalServerError{}, resp)
	})
	t.Run("Fails to get station status", func(t *testing.T) {
		stations, byMock := setup(t)
		infoResp := &clients.StationInfoResponse{
			LastUpdated: 111,
			Data: clients.StationsInfoData{
				Stations: []clients.StationInfo{
					{
						StationID:   "5431",
						StationName: "NRK",
						Latitude:    1.0,
						Longitude:   1.0,
						Address:     "NRK",
						Capacity:    21,
					},
				},
			},
		}
		byMock.EXPECT().GetStationInfo(ctx).Return(infoResp, nil).Times(1)
		byMock.EXPECT().GetStationStatus(ctx).Return(nil, fmt.Errorf("err")).Times(1)

		rawReq := httptest.NewRequest("GET", "/", nil).WithContext(ctx)
		params := operations.GetStationsInfoParams{
			HTTPRequest: rawReq,
		}
		resp := stations.Handle(params)
		assert.IsType(t, &operations.GetStationsInfoInternalServerError{}, resp)
	})
	t.Run("Success", func(t *testing.T) {
		stations, byMock := setup(t)
		infoResp := &clients.StationInfoResponse{
			LastUpdated: 111,
			Data: clients.StationsInfoData{
				Stations: []clients.StationInfo{
					{
						StationID:   "5431",
						StationName: "NRK",
						Latitude:    1.0,
						Longitude:   1.0,
						Address:     "NRK",
						Capacity:    21,
					},
				},
			},
		}
		statusResp := &clients.StationStatusResponse{
			LastUpdated: 111,
			Data: clients.StationStatusData{
				Stations: []clients.StationStatus{
					{
						StationID:         "5431",
						NumBikesAvailable: 3,
						NumDocksAvailable: 18,
						IsInstalled:       true,
						IsRenting:         true,
						IsReturning:       true,
						LastReported:      111,
					},
				},
			},
		}
		byMock.EXPECT().GetStationInfo(ctx).Return(infoResp, nil).Times(1)
		byMock.EXPECT().GetStationStatus(ctx).Return(statusResp, nil).Times(1)

		rawReq := httptest.NewRequest("GET", "/", nil).WithContext(ctx)
		params := operations.GetStationsInfoParams{
			HTTPRequest: rawReq,
		}
		resp := stations.Handle(params)
		assert.IsType(t, &operations.GetStationsInfoOK{}, resp)
		okResp := resp.(*operations.GetStationsInfoOK)
		assert.Equal(t, "5431", okResp.Payload.Stations[0].StationID)
		assert.Equal(t, "NRK", okResp.Payload.Stations[0].StationName)
		assert.Equal(t, int64(3), okResp.Payload.Stations[0].BikesAvailable)
		assert.Equal(t, int64(18), okResp.Payload.Stations[0].DocksAvailable)
	})
}

func setup(t *testing.T) (Stations, mocks.MockBySykkelClient) {
	ctrl := gomock.NewController(t)
	byClient := mocks.NewMockBySykkelClient(ctrl)
	return Stations{
		byClient: byClient,
	}, *byClient
}
