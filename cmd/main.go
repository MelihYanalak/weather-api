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
	var err error
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		return
	}
	defer logger.Log.Close()

	geoDb := adapter.NewTile38Repository("9851", "test_collection")
	geoDb.Initialize("new_york.geojson")
	weatherAPI := adapter.NewOpenWeatherAPI()
	cacheRedis := adapter.NewCacheRedis("6379")
	fmt.Println("program started")

	weatherService := application.NewWeatherService(geoDb, weatherAPI, cacheRedis)

	weatherController := controller.NewWeatherController(weatherService)

	router := mux.NewRouter()

	router.HandleFunc("/weather", weatherController.GetWeather).Methods("GET")

	http.ListenAndServe(":8080", router)
}
