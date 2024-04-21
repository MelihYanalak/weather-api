package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/redis/go-redis/v9"
)

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Geometry   json.RawMessage        `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type Tile38Repository struct {
	rdb *redis.Client
	key string
}

func (repo Tile38Repository) pointInPolygon(ctx context.Context, key string, point [2]float64) (bool, error) {
	cmd := fmt.Sprintf("WITHIN %s IDS POINT %f %f", key, point[1], point[0])
	result, err := repo.rdb.Do(ctx, "EVALSHA", cmd).Result()
	if err != nil {
		return false, err
	}

	ids, ok := result.([]interface{})
	if !ok || len(ids) == 0 {
		return false, nil // The point is not within any polygon
	}

	return true, nil // The point is inside at least one polygon
}

func (repo Tile38Repository) CheckLocation(latitude float64, longitude float64) (bool, error) {
	point := [2]float64{latitude, longitude}
	result, err := repo.pointInPolygon(context.TODO(), repo.key, point)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (repo Tile38Repository) Initialize(filePath string, port string, collectionName string) error {
	//filePath := "build/new_york.geojson"
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
	repo.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port, // or the address of your Tile38 server //PORT : ******9851*********
		Password: "",                  // no password set
		DB:       0,                   // use default DB
	})
	ctx := context.Background()
	repo.rdb.Do(ctx, "DROP", collectionName).Result()
	// Insert each feature into the Tile38 database
	for i, feature := range geoJSON.Features {
		id := fmt.Sprintf("feature_%d", i) // Generate a unique ID for each feature
		geoJSONStr, err := feature.Geometry.MarshalJSON()
		if err != nil {
			log.Printf("Error marshalling geometry: %v", err)
			continue
		}

		_, err = repo.rdb.Do(ctx, "SET", collectionName, id, "OBJECT", string(geoJSONStr)).Result()
		if err != nil {
			log.Printf("Error inserting feature %d: %v", i, err)
			continue
		}
	}

	fmt.Println("GeoJSON data has been loaded into Tile38")
	return nil
}
