package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// CheckWeatherHandler handles the "/weather" endpoint.
func (c *WeatherController) CheckWeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")

	// Convert lat and long to float64 (you might want to add error handling here)
	latFloat, _ := strconv.ParseFloat(lat, 64)
	longFloat, _ := strconv.ParseFloat(long, 64)

	// Call the weather service method
	weatherData, err := c.weatherService.CheckWeather(latFloat, longFloat)
	if err != nil {
		// Handle error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Encode weather data to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
