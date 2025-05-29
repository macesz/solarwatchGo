package geolocation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/macesz/solarwatchGo/internal/domain"
)

type Service struct {
	apiId string
}

func NewService(apiId string) *Service {
	return &Service{apiId: apiId}
}

func (s Service) GetReport(city, countryCode, stateCode string) (*domain.GeoLocation, error) {
	combinedLocation := fmt.Sprintf("%s,%s", city, countryCode)
	if stateCode != "" {
		combinedLocation = combinedLocation + "," + stateCode
	}

	api := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", combinedLocation, s.apiId)

	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rsb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loc []geoLocDto

	err = json.Unmarshal(rsb, &loc)
	if err != nil {
		return nil, err
	}

	if len(loc) == 0 {
		return nil, nil
	}

	return &domain.GeoLocation{
		Lat: loc[0].Lat,
		Lon: loc[0].Lon,
	}, nil
}
