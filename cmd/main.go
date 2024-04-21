package main

import (
	"fmt"
	"net/http"

	"github.com/MelihYanalak/weather-api/internal/application"
	"github.com/MelihYanalak/weather-api/internal/controller"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("program started")
	// Initialize dependencies
	weatherService := application.NewWeatherService()

	// Create a new WeatherController instance
	weatherController := controller.NewWeatherController(weatherService)

	// Create a new HTTP router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/weather", weatherController.CheckWeatherHandler).Methods("GET")

	// Start the HTTP server
	http.ListenAndServe(":8080", router)
}
