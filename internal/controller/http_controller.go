package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MelihYanalak/weather-api/internal/domain"
)

type WeatherService interface {
	GetWeather(ctx context.Context, lat, long float64) (domain.Weather, error)
}
type WeatherController struct {
	weatherService WeatherService
}

func NewWeatherController(weatherService WeatherService) *WeatherController {
	return &WeatherController{weatherService: weatherService}
}

type Location struct {
	Lat  float64
	Long float64
}

func (c *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {

	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weatherData, err := c.weatherService.GetWeather(context.TODO(), location.Lat, location.Long)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		WeatherDTO{
			Definition:  weatherData.Definition,
			Description: weatherData.Description,
			Temperature: weatherData.Temperature,
			Humidity:    weatherData.Humidity,
		},
	)
}
