package main

import (
	"net/http"
	"os"

	"log"

	"github.com/MelihYanalak/weather-api/internal/adapter"
	"github.com/MelihYanalak/weather-api/internal/application"
	"github.com/MelihYanalak/weather-api/internal/controller"
	"github.com/MelihYanalak/weather-api/internal/logger"
	"github.com/gorilla/mux"
)

func main() {
	logger, err := logger.NewFileLogger(logger.DebugLevel, "weather-api.log")
	if err != nil {
		log.Fatal("Could not create logger")
	}
	defer logger.Close()

	tile38Host := os.Getenv("TILE38_HOST")
	redisHost := os.Getenv("REDIS_HOST")
	owmKey := os.Getenv("OWM_API_KEY")

	geoDb := adapter.NewTile38Repository(tile38Host)
	cacheRedis := adapter.NewCacheRedis(redisHost)
	weatherAPI := adapter.NewOpenWeatherAPI(owmKey)

	weatherService := application.NewWeatherService(geoDb, weatherAPI, cacheRedis, logger)
	weatherController := controller.NewWeatherController(weatherService)

	router := mux.NewRouter()

	router.HandleFunc("/weather", weatherController.GetWeather).Methods("GET")

	logger.Info("Program started")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Error("Could not start server")
	}
}
