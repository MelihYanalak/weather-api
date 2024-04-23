package repository

import (
	"context"

	"github.com/MelihYanalak/weather-api/internal/domain"
)

type WeatherAPI interface {
	Get(ctx context.Context, float64, longitude float64) (domain.Weather, error)
}
