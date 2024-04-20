package application

import (
	"github.com/MelihYanalak/weather-api/internal/domain"
)

type IWeatherService interface {
	CheckWeather(lat, long float64) (domain.Weather, error)
}
