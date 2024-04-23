package controller

import (
	"encoding/json"
	"net/http"

	"github.com/MelihYanalak/weather-api/internal/application"
)

type WeatherController struct {
	weatherService application.IWeatherService
}

func NewWeatherController(weatherService application.IWeatherService) *WeatherController {
	return &WeatherController{weatherService: weatherService}
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func (c *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	weatherData, err := c.weatherService.GetWeather(location.Lat, location.Long)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
