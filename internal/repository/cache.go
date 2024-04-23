package repository

import "github.com/MelihYanalak/weather-api/internal/domain"

type ICacheRepository interface {
	RetrieveData(key string) (domain.Weather, error)
	InsertData(key string, weather domain.Weather) error
}
