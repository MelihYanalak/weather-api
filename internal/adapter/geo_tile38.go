package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/MelihYanalak/weather-api/internal/logger"
	"github.com/redis/go-redis/v9"
)

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   Geometry               `json:"geometry"`
}

type Geometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}

type Tile38Repository struct {
	rdb *redis.Client
	key string
}

func NewTile38Repository(port string, collectionName string) *Tile38Repository {
	rdb := redis.NewClient(&redis.Options{
		Addr: "tile38:" + port,
	})
	return &Tile38Repository{
		rdb: rdb,
		key: collectionName,
	}
}

func (repo Tile38Repository) CheckLocation(latitude float64, longitude float64) (bool, error) {
	result, err := repo.rdb.Do(context.Background(), "INTERSECTS", repo.key, "IDS", "POINT", fmt.Sprintf("%f", latitude), fmt.Sprintf("%f", longitude)).Result()
	if err != nil {
		return false, err
	}
	fmt.Println(result)
	resultSlice, ok := result.([]interface{})
	if !ok {
		return false, fmt.Errorf("unexpected result format")
	}

	if len(resultSlice) < 2 {
		return false, nil
	}

	idsSlice, ok := resultSlice[1].([]interface{})
	if !ok || len(idsSlice) == 0 {
		return false, nil
	}

	return true, nil
}

func (repo Tile38Repository) Initialize(filePath string) error {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var geoJSON FeatureCollection
	if err := json.Unmarshal(data, &geoJSON); err != nil {
		log.Fatalf("Error parsing GeoJSON: %v", err)
	}

	ctx := context.Background()
	_, err = repo.rdb.Do(ctx, "DROP", repo.key).Result()
	if err != nil {
		log.Printf("Error dropping collection: %v", err)
	}

	for idx, feature := range geoJSON.Features {
		id := fmt.Sprintf("feature_%d", idx)
		geoJSONStr, err := json.Marshal(feature.Geometry)
		if err != nil {
			log.Printf("Error marshalling geometry: %v", err)
			continue
		}

		_, err = repo.rdb.Do(ctx, "SET", repo.key, id, "OBJECT", string(geoJSONStr)).Result()
		if err != nil {
			logger.Log.Error("Error inserting feature" + id + err.Error())
		}
	}

	logger.Log.Info("GeoJSON data has been loaded into Tile38")
	return nil
}
