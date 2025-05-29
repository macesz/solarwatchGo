package solarwatch

type SolarWatchDTO struct {
	Status  string `json:"status"`
	Results struct {
		Sunrise string `json:"sunrise"`
		Sunset  string `json:"sunset"`
	} `json:"results"`
}
