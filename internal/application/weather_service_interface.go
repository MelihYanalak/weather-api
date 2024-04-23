package application

import (
	"github.com/MelihYanalak/weather-api/internal/domain"
)

type IWeatherService interface {
	GetWeather(lat, long float64) (domain.Weather, error)
}
