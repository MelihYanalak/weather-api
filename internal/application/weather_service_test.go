package application

import (
	"context"
	"errors"
	"testing"

	"github.com/MelihYanalak/weather-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

// Mock implementations of interfaces
type mockGeoRepo struct {
	locationExists bool
	err            error
}

func (m *mockGeoRepo) CheckLocation(ctx context.Context, latitude, longitude float64) (bool, error) {
	return m.locationExists, m.err
}

func (m *mockGeoRepo) Initialize(ctx context.Context, filePath string) error {
	return nil
}

type mockWeatherAPI struct {
	weather domain.Weather
	err     error
}

func (m *mockWeatherAPI) Get(ctx context.Context, lat, long float64) (domain.Weather, error) {
	return m.weather, m.err
}

type mockCacheRepo struct {
	weather  domain.Weather
	indexErr error
	getErr   error
	setErr   error
}

func (m *mockCacheRepo) Get(ctx context.Context, key string) (domain.Weather, error) {
	return m.weather, m.getErr
}

func (m *mockCacheRepo) Set(ctx context.Context, key string, weather domain.Weather) error {
	return m.setErr
}

func (m *mockCacheRepo) IndexKey(ctx context.Context, lat, long float64) (string, error) {
	return "key", m.indexErr
}

type mockLogger struct{}

func (m *mockLogger) Error(message string)             {}
func (m *mockLogger) Debug(message string)             {}
func (m *mockLogger) Info(message string)              {}
func (m *mockLogger) Warning(message string)           {}
func (m *mockLogger) log(level string, message string) {}

// Tests
func TestGetWeather(t *testing.T) {
	logger := &mockLogger{}
	tests := []struct {
		name            string
		geoRepo         *mockGeoRepo
		weatherApi      *mockWeatherAPI
		cacheRepo       *mockCacheRepo
		expectedWeather domain.Weather
		expectedError   error
	}{
		{
			name:            "API fetch and cache set successful",
			geoRepo:         &mockGeoRepo{locationExists: true},
			weatherApi:      &mockWeatherAPI{weather: domain.Weather{Temperature: 25}},
			cacheRepo:       &mockCacheRepo{getErr: errors.New("cache miss"), setErr: nil},
			expectedWeather: domain.Weather{Temperature: 25},
			expectedError:   nil,
		},
		{
			name:            "API fetch failure after cache miss",
			geoRepo:         &mockGeoRepo{locationExists: true},
			weatherApi:      &mockWeatherAPI{err: errors.New("api failure")},
			cacheRepo:       &mockCacheRepo{getErr: errors.New("cache miss")},
			expectedWeather: domain.Weather{},
			expectedError:   errors.New("api failure"),
		},
		{
			name:            "Location check fails",
			geoRepo:         &mockGeoRepo{err: errors.New("location error")},
			weatherApi:      &mockWeatherAPI{},
			cacheRepo:       &mockCacheRepo{},
			expectedWeather: domain.Weather{},
			expectedError:   errors.New("location error"),
		},
		{
			name:            "Location not in the market region",
			geoRepo:         &mockGeoRepo{locationExists: false},
			weatherApi:      &mockWeatherAPI{},
			cacheRepo:       &mockCacheRepo{},
			expectedWeather: domain.Weather{},
			expectedError:   errors.New("point not in the market region"),
		},
		{
			name:            "Index key failure",
			geoRepo:         &mockGeoRepo{locationExists: true},
			weatherApi:      &mockWeatherAPI{},
			cacheRepo:       &mockCacheRepo{indexErr: errors.New("key generation error")},
			expectedWeather: domain.Weather{},
			expectedError:   errors.New("key generation error"),
		},
		{
			name:            "Cache hit",
			geoRepo:         &mockGeoRepo{locationExists: true},
			weatherApi:      &mockWeatherAPI{},
			cacheRepo:       &mockCacheRepo{weather: domain.Weather{Temperature: 22}},
			expectedWeather: domain.Weather{Temperature: 22},
			expectedError:   nil,
		},
		{
			name:            "Cache set failure",
			geoRepo:         &mockGeoRepo{locationExists: true},
			weatherApi:      &mockWeatherAPI{weather: domain.Weather{Temperature: 23}},
			cacheRepo:       &mockCacheRepo{getErr: errors.New("cache miss"), setErr: errors.New("cache set error")},
			expectedWeather: domain.Weather{Temperature: 23},
			expectedError:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewWeatherService(test.geoRepo, test.weatherApi, test.cacheRepo, logger)
			weather, err := service.GetWeather(context.Background(), 0, 0)
			assert.Equal(t, test.expectedWeather, weather)
			if test.expectedError != nil {
				assert.EqualError(t, err, test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
