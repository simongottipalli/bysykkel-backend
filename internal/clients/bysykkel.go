package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BysykkelClient struct {
	client  http.Client
	baseUrl url.URL
}

func NewBysykkelClient() (*BysykkelClient, error) {
	baseUrl, err := url.Parse("https://gbfs.urbansharing.com/oslobysykkel.no/")
	if err != nil {
		return nil, err
	}
	httpClient := http.DefaultClient

	return &BysykkelClient{
		client:  *httpClient,
		baseUrl: *baseUrl,
	}, nil
}

func (b *BysykkelClient) GetStationInfo(ctx context.Context) (*StationInfoResponse, error) {
	path := "station_information.json"
	resp, err := b.makeRequest(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to make station_information request: %w", err)
	}

	var responseData StationInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("failed to decode response from API: %w", err)
	}
	return &responseData, nil
}

func (b *BysykkelClient) GetStationStatus(ctx context.Context) (*StationStatusResponse, error) {
	path := "station_status.json"

	resp, err := b.makeRequest(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to make station_status request: %w", err)
	}

	var responseData StationStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("failed to decode response from API: %w", err)
	}
	return &responseData, nil
}

func (b *BysykkelClient) makeRequest(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.baseUrl.JoinPath(path).String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Client-Identifier", "origo/intervju")

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return resp, nil
}
