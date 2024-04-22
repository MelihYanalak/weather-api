package controller

import (
	"encoding/json"
	"net/http"

	"github.com/MelihYanalak/weather-api/internal/application"
)

// WeatherController handles HTTP requests related to weather.
type WeatherController struct {
	weatherService application.IWeatherService
}

// NewWeatherController creates a new WeatherController instance.
func NewWeatherController(weatherService application.IWeatherService) *WeatherController {
	return &WeatherController{weatherService: weatherService}
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// CheckWeatherHandler handles the "/weather" endpoint.
func (c *WeatherController) CheckWeatherHandler(w http.ResponseWriter, r *http.Request) {
	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// You can directly access lat and long as float64
	lat := location.Lat
	long := location.Long

	// Call the weather service method
	weatherData, err := c.weatherService.CheckWeather(lat, long)
	if err != nil {
		// Handle error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Encode weather data to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
