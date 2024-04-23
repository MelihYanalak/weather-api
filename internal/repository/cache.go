package repository

import (
	"context"

	"github.com/MelihYanalak/weather-api/internal/domain"
)

type CacheRepository interface {
	Get(ctx context.Context, key string) (domain.Weather, error)
	Set(ctx context.Context, key string, weather domain.Weather) error
	IndexKey(ctx context.Context, lat, long float64) (string, error)
}
