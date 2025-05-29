package solarwatch

import (
	"time"

	"github.com/macesz/solarwatchGo/internal/domain"
)

type GeoLocationGetter interface {
	GetReport(city, countryCode, stateCode string) (*domain.GeoLocation, error)
}

type SolarWatchGetter interface {
	GetReport(lat float64, lon float64, date time.Time) (*domain.SolarWatch, error)
}
