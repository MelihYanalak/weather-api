package application

import (
	"context"
	"fmt"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/MelihYanalak/weather-api/internal/logger"
	"github.com/MelihYanalak/weather-api/internal/repository"
)

type CacheRepository interface {
	Get(ctx context.Context, key string) (domain.Weather, error)
	Set(ctx context.Context, key string, weather domain.Weather) error
	IndexKey(ctx context.Context, lat, long float64) (string, error)
}

type GeoRepository interface {
	CheckLocation(ctx context.Context, latitude float64, longitude float64) (bool, error)
	Initialize(ctx context.Context, filePath string) error
}

type WeatherAPI interface {
	Get(ctx context.Context, float64, longitude float64) (domain.Weather, error)
}

type WeatherService struct {
	geoRepo    repository.GeoRepository
	weatherApi repository.WeatherAPI
	cacheRepo  repository.CacheRepository
	logger     logger.Logger
}

func NewWeatherService(geoRepository repository.GeoRepository, weatherApi repository.WeatherAPI, cacheRepository repository.CacheRepository, logger logger.Logger) *WeatherService {
	return &WeatherService{
		geoRepo:    geoRepository,
		weatherApi: weatherApi,
		cacheRepo:  cacheRepository,
		logger:     logger,
	}
}
func (ws WeatherService) GetWeather(ctx context.Context, lat, long float64) (domain.Weather, error) {
	result, err := ws.geoRepo.CheckLocation(ctx, lat, long)
	if err != nil {
		return domain.Weather{}, err
	}
	if !result {
		ws.logger.Error("Point not in market region")
		//define specific err type for it
		return domain.Weather{}, err
	}

	key, _ := ws.cacheRepo.IndexKey(ctx, lat, long)
	weather, err := ws.cacheRepo.Get(ctx, key)
	if err != nil { //TODO : convert err to not found type
		fmt.Println("Not found in cache, requesting to API")
		weather, err = ws.weatherApi.Get(ctx, lat, long)
		if err != nil {
			return domain.Weather{}, err
		}
		ws.cacheRepo.Set(ctx, key, weather)

	} else {
		fmt.Println("found in cache")
	}
	return weather, nil

}
