package adapter

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Tile38Repository struct {
	rdb *redis.Client
	key string
}

func NewTile38Repository(host string) *Tile38Repository {
	rdb := redis.NewClient(&redis.Options{
		Addr: host,
	})
	return &Tile38Repository{
		rdb: rdb,
		key: "test_collection",
	}
}

func (repo Tile38Repository) CheckLocation(ctx context.Context, latitude float64, longitude float64) (bool, error) {
	result, err := repo.rdb.Do(ctx, "INTERSECTS", repo.key, "IDS", "POINT", fmt.Sprintf("%f", latitude), fmt.Sprintf("%f", longitude)).Result()
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
