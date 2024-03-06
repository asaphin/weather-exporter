package app

import (
	"github.com/asaphin/weather-exporter/metrics"
	"github.com/rs/zerolog/log"
)

type WeatherService struct {
	lat    float64
	lon    float64
	client WeatherClient
}

func (w *WeatherService) UpdateWeatherData() {
	weather, err := w.client.GetWeather(w.lat, w.lon)
	if err != nil {
		log.Err(err).Msg("unable to get weather data")
	}

	metrics.RegisterTemperature(weather.Temperature, weather.LocationName)
	metrics.RegisterPressure(weather.Pressure, weather.LocationName)
	metrics.RegisterHumidity(weather.Humidity, weather.LocationName)
	metrics.RegisterWindSpeed(weather.WindSpeed, weather.LocationName)
	metrics.RegisterWindDirection(weather.WindDirection, weather.LocationName)
	metrics.RegisterCloudsAll(weather.CloudsAll, weather.LocationName)
	metrics.RegisterRainOneHour(weather.RainOneHour, weather.LocationName)
}

func NewWeatherService(client WeatherClient, lat, lon float64) *WeatherService {
	return &WeatherService{
		lat:    lat,
		lon:    lon,
		client: client,
	}
}
