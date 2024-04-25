package adapter

import (
	"context"
	"fmt"

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
		fmt.Println("can not create new Current")
	}
	err = currentWeather.CurrentByCoordinates(&owm.Coordinates{Longitude: longitude, Latitude: latitude})
	if err != nil {
		fmt.Println("can not get coordinates:", latitude, "-", longitude)
	}
	return domain.Weather{
		Definition:  currentWeather.Weather[0].Main,
		Description: currentWeather.Weather[0].Description,
		Temperature: currentWeather.Main.Temp,
		Humidity:    currentWeather.Main.Humidity,
	}, nil
}
