package application

import (
	"fmt"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/MelihYanalak/weather-api/internal/repository"
)

type WeatherService struct {
	gr repository.GeoRepository
}

func NewWeatherService(geo_repo repository.GeoRepository) *WeatherService {
	return &WeatherService{
		gr: geo_repo,
	}
}
func (ws WeatherService) CheckWeather(lat, long float64) (domain.Weather, error) {
	result, err := ws.gr.CheckLocation(lat, long)
	if err != nil {
		fmt.Println("error")
	}
	weatherData := domain.Weather{
		Temperature: 15,
		Condition:   "good",
		Valid:       result,
	}
	return weatherData, nil

}
