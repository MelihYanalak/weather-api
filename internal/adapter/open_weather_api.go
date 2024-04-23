package adapter

import (
	"context"

	"github.com/MelihYanalak/weather-api/internal/domain"
	owm "github.com/briandowns/openweathermap"
)

var API_KEY string = "98f280d8961dcbc064b1d69f980c5c5a"

type OpenWeatherAPI struct {
}

func NewOpenWeatherAPI() *OpenWeatherAPI {
	return &OpenWeatherAPI{}
}

func (owApi OpenWeatherAPI) Get(ctx context.Context, latitude float64, longitude float64) (domain.Weather, error) {

	currentWeather, err := owm.NewCurrent("C", "en", API_KEY)
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
