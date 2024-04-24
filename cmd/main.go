package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/MelihYanalak/weather-api/internal/adapter"
	"github.com/MelihYanalak/weather-api/internal/application"
	"github.com/MelihYanalak/weather-api/internal/controller"
	"github.com/MelihYanalak/weather-api/internal/logger"
	"github.com/gorilla/mux"
)

func main() {
	defer logger.Log.Close()
	tile38Host := os.Getenv("TILE38_HOST")
	redisHost := os.Getenv("REDIS_HOST")
	owmKey := os.Getenv("OWM_API_KEY")
	geoDb := adapter.NewTile38Repository(tile38Host)
	cacheRedis := adapter.NewCacheRedis(redisHost)
	weatherAPI := adapter.NewOpenWeatherAPI(owmKey)

	geoDb.Initialize(context.Background(), "new_york.geojson")
	fmt.Println("program started")

	weatherService := application.NewWeatherService(geoDb, weatherAPI, cacheRedis)
	weatherController := controller.NewWeatherController(weatherService)

	router := mux.NewRouter()

	router.HandleFunc("/weather", weatherController.GetWeather).Methods("GET")

	http.ListenAndServe(":8080", router)
}
