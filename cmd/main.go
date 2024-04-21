package main

import (
	"fmt"
	"net/http"

	"github.com/MelihYanalak/weather-api/internal/adapter"
	"github.com/MelihYanalak/weather-api/internal/application"
	"github.com/MelihYanalak/weather-api/internal/controller"
	"github.com/MelihYanalak/weather-api/internal/logger"
	"github.com/gorilla/mux"
)

func main() {

	logger, err := logger.NewFileLogger(logger.DebugLevel, "weather-api.log")
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		return
	}
	defer logger.Close()

	geo_db := adapter.NewTile38Repository("9851", "test_collection")
	geo_db.Initialize("build/new_york.geojson")

	fmt.Println("program started")
	// Initialize dependencies

	weatherService := application.NewWeatherService(geo_db)

	// Create a new WeatherController instance
	weatherController := controller.NewWeatherController(weatherService)

	// Create a new HTTP router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/weather", weatherController.CheckWeatherHandler).Methods("GET")

	// Start the HTTP server
	http.ListenAndServe(":8080", router)
}
