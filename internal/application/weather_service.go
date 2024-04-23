package application

import (
	"context"
	"fmt"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/MelihYanalak/weather-api/internal/repository"
)

type WeatherService struct {
	geoRepo    repository.GeoRepository
	weatherApi repository.WeatherAPI
	cacheRepo  repository.CacheRepository
}

func NewWeatherService(geoRepository repository.GeoRepository, weatherApi repository.WeatherAPI, cacheRepository repository.CacheRepository) *WeatherService {
	return &WeatherService{
		geoRepo:    geoRepository,
		weatherApi: weatherApi,
		cacheRepo:  cacheRepository,
	}
}
func (ws WeatherService) GetWeather(ctx context.Context, lat, long float64) (domain.Weather, error) {
	result, err := ws.geoRepo.CheckLocation(ctx, lat, long)
	if err != nil {
		return domain.Weather{}, err
	}
	if !result {
		fmt.Println("point not in market region")
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
