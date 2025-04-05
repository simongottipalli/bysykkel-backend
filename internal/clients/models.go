package clients

type StationInfoResponse struct {
	LastUpdated int              `json:"last_updated"`
	Data        StationsInfoData `json:"data"`
}

type StationsInfoData struct {
	Stations []StationInfo `json:"stations"`
}

type StationInfo struct {
	StationID   string  `json:"station_id"`
	StationName string  `json:"name"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Capacity    int     `json:"capacity"`
}

type StationStatusResponse struct {
	LastUpdated int               `json:"last_updated"`
	Data        StationStatusData `json:"data"`
}

type StationStatusData struct {
	Stations []StationStatus `json:"stations"`
}

type StationStatus struct {
	IsInstalled       bool   `json:"is_installed"`
	IsRenting         bool   `json:"is_renting"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumDocksAvailable int    `json:"num_docks_available"`
	LastReported      int    `json:"last_reported"`
	IsReturning       bool   `json:"is_returning"`
	StationID         string `json:"station_id"`
}
