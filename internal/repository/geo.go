package repository

import "github.com/MelihYanalak/weather-api/internal/domain"

type GeoRepository interface {
	RetrieveData(key string) (domain.Weather, error)
	InsertData(key string, weather domain.Weather) error
}
