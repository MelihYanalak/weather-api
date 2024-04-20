package application

import "github.com/MelihYanalak/weather-api/internal/domain"

type WeatherService struct {
	//core object

}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}
func (ws WeatherService) CheckWeather(lat, long float64) (domain.Weather, error) {

	weatherData := domain.Weather{
		Temperature: 15,
		Condition:   "good",
	}
	return weatherData, nil

}
