package repository

import "github.com/MelihYanalak/weather-api/internal/domain"

type IWeatherAPI interface {
	GetWeatherData(latitude float64, longitude float64) (domain.Weather, error)
}
