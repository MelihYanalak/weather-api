package application

import (
	"fmt"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/MelihYanalak/weather-api/internal/repository"
)

type WeatherService struct {
	gr   repository.IGeoRepository
	wApi repository.IWeatherAPI
	cr   repository.ICacheRepository
}

func NewWeatherService(geoRepository repository.IGeoRepository, weatherApi repository.IWeatherAPI, cacheRepository repository.ICacheRepository) *WeatherService {
	return &WeatherService{
		gr:   geoRepository,
		wApi: weatherApi,
		cr:   cacheRepository,
	}
}
func (ws WeatherService) GetWeather(lat, long float64) (domain.Weather, error) {
	result, err := ws.gr.CheckLocation(lat, long)
	if err != nil {
		return domain.Weather{}, err
	}
	if !result {
		fmt.Println("point not in market region")
		//define specific err type for it
		return domain.Weather{}, err
	}

	key, _ := ws.cr.IndexKey(lat, long)
	weather, err := ws.cr.RetrieveData(key)
	if err != nil { //TODO : convert err to not found type
		fmt.Println("Not found in cache, requesting to API")
		weather, err = ws.wApi.GetWeatherData(lat, long)
		if err != nil {
			return domain.Weather{}, err
		}
		ws.cr.InsertData(key, weather)

	} else {
		fmt.Println("found in cache")
	}
	return weather, nil

}
