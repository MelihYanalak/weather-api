package application

import (
	"fmt"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/MelihYanalak/weather-api/internal/repository"
)

type WeatherService struct {
	gr   repository.IGeoRepository
	wApi repository.IWeatherAPI
}

func NewWeatherService(geoRepository repository.IGeoRepository, weatherApi repository.IWeatherAPI) *WeatherService {
	return &WeatherService{
		gr:   geoRepository,
		wApi: weatherApi,
	}
}
func (ws WeatherService) GetWeather(lat, long float64) (domain.Weather, error) {
	result, err := ws.gr.CheckLocation(lat, long)
	if err != nil {
		return domain.Weather{}, err
	}
	if !result {
		fmt.Println("point not in market region")
		//define specific err type for it
		return domain.Weather{}, err
	}

	weather, err := ws.wApi.GetWeatherData(lat, long)
	if err != nil {
		return domain.Weather{}, err
	}
	return weather, nil

}
