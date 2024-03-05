package app

import "github.com/asaphin/weather-exporter/domain"

type WeatherClient interface {
	GetWeather(lat, lon float64) (*domain.Weather, error)
}
