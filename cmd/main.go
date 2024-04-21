package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/MelihYanalak/weather-api/internal/application"
	"github.com/MelihYanalak/weather-api/internal/controller"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Geometry   json.RawMessage        `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("saa:", dir)
	// Print the current directory
	fmt.Println("Current Directory:", dir)
	filePath := "build/new_york.geojson"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Parse the GeoJSON data
	var geoJSON FeatureCollection
	if err := json.Unmarshal(data, &geoJSON); err != nil {
		log.Fatalf("Error parsing GeoJSON: %v", err)
	}

	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:9851", // or the address of your Tile38 server
	})
	ctx := context.Background()
	rdb.Do(ctx, "DROP", "your_collection").Result()
	// Insert each feature into the Tile38 database
	for i, feature := range geoJSON.Features {
		id := fmt.Sprintf("feature_%d", i) // Generate a unique ID for each feature
		geoJSONStr, err := feature.Geometry.MarshalJSON()
		if err != nil {
			log.Printf("Error marshalling geometry: %v", err)
			continue
		}

		_, err = rdb.Do(ctx, "SET", "your_collection", id, "OBJECT", string(geoJSONStr)).Result()
		if err != nil {
			log.Printf("Error inserting feature %d: %v", i, err)
			continue
		}
	}

	fmt.Println("GeoJSON data has been loaded into Tile38")

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
