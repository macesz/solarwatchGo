package solarwatch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/macesz/solarwatchGo/internal/domain"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) GetReport(lat, lon float64, date time.Time) (*domain.SolarWatch, error) {
	api := fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%f&lng=%f&date=%s", lat, lon, date.Format(time.DateOnly))
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	// defer is executed when exit the func
	defer resp.Body.Close()

	rsb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response SolarWatchDTO
	err = json.Unmarshal(rsb, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "OK" {
		return nil, fmt.Errorf("eoor")
	}

	return &domain.SolarWatch{
		Sunrise: response.Results.Sunrise,
		Sunset:  response.Results.Sunset,
	}, nil

}
