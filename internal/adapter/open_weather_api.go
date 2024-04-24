package adapter

import (
	"context"

	"github.com/MelihYanalak/weather-api/internal/domain"
	owm "github.com/briandowns/openweathermap"
)

type OpenWeatherAPI struct {
	key string
}

func NewOpenWeatherAPI(apiKey string) *OpenWeatherAPI {
	return &OpenWeatherAPI{
		key: apiKey,
	}
}

func (owApi OpenWeatherAPI) Get(ctx context.Context, latitude float64, longitude float64) (domain.Weather, error) {

	currentWeather, err := owm.NewCurrent("C", "en", owApi.key)
	if err != nil {
		//error and exit
	}
	err = currentWeather.CurrentByCoordinates(&owm.Coordinates{Longitude: longitude, Latitude: latitude})

	return domain.Weather{
		Definition:  currentWeather.Weather[0].Main,
		Description: currentWeather.Weather[0].Description,
		Temperature: currentWeather.Main.Temp,
		Humidity:    currentWeather.Main.Humidity,
	}, nil
}
