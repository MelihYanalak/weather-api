package controller

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MelihYanalak/weather-api/internal/domain"

	"github.com/stretchr/testify/assert"
)

type mockWeatherService struct {
	weather domain.Weather
	err     error
}

func (m *mockWeatherService) GetWeather(ctx context.Context, lat, long float64) (domain.Weather, error) {
	return m.weather, m.err
}

func TestWeatherController_GetWeather(t *testing.T) {
	// Sample weather data and error for use in tests
	sampleWeather := domain.Weather{
		Definition:  "Sunny",
		Description: "Clear sky",
		Temperature: 25,
		Humidity:    70,
	}

	tests := []struct {
		name         string
		body         string
		weather      domain.Weather
		serviceError error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "successful weather retrieval",
			body:         `{"Lat": 34.05, "Long": -118.25}`,
			weather:      sampleWeather,
			serviceError: nil,
			expectedCode: http.StatusOK,
			expectedBody: "{\"definition\":\"Sunny\",\"description\":\"Clear sky\",\"temperature\":25,\"humidity\":70}\n",
		},
		{
			name:         "invalid request body",
			body:         `{"Lat": "bad input", "Long": "no number"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "json: cannot unmarshal string into Go struct field Location.Lat of type float64\n",
		},
		{
			name:         "service error",
			body:         `{"Lat": 34.05, "Long": -118.25}`,
			serviceError: errors.New("service error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: "service error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/weather", bytes.NewBufferString(tt.body))
			rr := httptest.NewRecorder()

			mockService := &mockWeatherService{
				weather: tt.weather,
				err:     tt.serviceError,
			}
			controller := NewWeatherController(mockService)

			handler := http.HandlerFunc(controller.GetWeather)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
