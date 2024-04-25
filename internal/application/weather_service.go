package application

import (
	"context"
	"errors"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/MelihYanalak/weather-api/internal/logger"
	"github.com/MelihYanalak/weather-api/internal/repository"
)

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
		ws.logger.Error(err.Error())
		return domain.Weather{}, err
	}
	if !result {
		ws.logger.Error("point not in the market region")
		return domain.Weather{}, errors.New("point not in the market region")
	}

	key, err := ws.cacheRepo.IndexKey(ctx, lat, long)
	if err != nil {
		ws.logger.Error(err.Error())
		return domain.Weather{}, err
	}
	weather, err := ws.cacheRepo.Get(ctx, key)
	if err != nil {
		weather, err = ws.weatherApi.Get(ctx, lat, long)
		if err != nil {
			ws.logger.Error(err.Error())
			return domain.Weather{}, err
		}
		ws.cacheRepo.Set(ctx, key, weather)
	}

	return weather, nil

}
