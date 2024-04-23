package adapter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/redis/go-redis/v9"
	"github.com/uber/h3-go/v4"
)

type CacheRedis struct {
	client *redis.Client
}

func NewCacheRedis(portNumber string) *CacheRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:" + portNumber, // Redis server address
		Password: "",                    // no password set
		DB:       0,                     // use default DB
	})
	return &CacheRedis{client: rdb}
}

func (c *CacheRedis) RetrieveData(key string) (domain.Weather, error) {
	val, err := c.client.Get(context.TODO(), key).Result()
	if err != nil {
		return domain.Weather{}, err
	}

	var weatherData domain.Weather
	err = json.Unmarshal([]byte(val), &weatherData)
	if err != nil {
		return domain.Weather{}, err
	}

	return weatherData, nil
}

func (c *CacheRedis) InsertData(key string, weather domain.Weather) error {
	data, err := json.Marshal(weather)
	if err != nil {
		return err
	}

	err = c.client.Set(context.Background(), key, string(data), 0).Err()
	if err != nil {
		return err
	}

	return nil
}
func (c *CacheRedis) IndexKey(lat, long float64) (string, error) {
	h3Index := h3.LatLngToCell(h3.LatLng{Lat: lat, Lng: long}, 8)
	h3IndexStr := fmt.Sprintf("%x", h3Index)
	return h3IndexStr, nil
}
