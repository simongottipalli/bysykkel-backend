package handlers

import (
	"bysykkel/internal/api/models"
	"bysykkel/internal/api/restapi/operations"
	"bysykkel/internal/clients"
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type BySykkelClient interface {
	GetStationInfo(ctx context.Context) (*clients.StationInfoResponse, error)
	GetStationStatus(ctx context.Context) (*clients.StationStatusResponse, error)
}

type Stations struct {
	byClient BySykkelClient
}

func NewStations(byClient BySykkelClient) *Stations {
	return &Stations{
		byClient: byClient,
	}
}

func (s Stations) Handle(params operations.GetStationsInfoParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	stationInfo, err := s.byClient.GetStationInfo(ctx)
	if err != nil {
		fmt.Print(err)
		return operations.NewGetStationsInfoInternalServerError().WithPayload(&models.Error{
			Code:    "GetStationsError",
			Message: "Failed to get stations",
		})
	}

	stationStatus, err := s.byClient.GetStationStatus(ctx)
	if err != nil {
		fmt.Print(err)
		return operations.NewGetStationsInfoInternalServerError().WithPayload(&models.Error{
			Code:    "GetStationsError",
			Message: "Failed to get stations",
		})
	}

	resp := s.getStationInfo(*stationInfo, *stationStatus)

	return operations.NewGetStationsInfoOK().WithPayload(&models.StationStatusResponse{
		LastUpdated: 111,
		Stations:    resp,
	})
}

func (s Stations) getStationInfo(stationInfo clients.StationInfoResponse, statusStatus clients.StationStatusResponse) []*models.StationInfo {
	var respStationInfo []*models.StationInfo
	for _, info := range stationInfo.Data.Stations {
		si := models.StationInfo{
			StationID:   info.StationID,
			StationName: info.StationName,
		}
		for _, status := range statusStatus.Data.Stations {
			if status.StationID == info.StationID {
				si.BikesAvailable = int64(status.NumBikesAvailable)
				si.DocksAvailable = int64(status.NumDocksAvailable)
			}
		}
		respStationInfo = append(respStationInfo, &si)
	}
	return respStationInfo
}
