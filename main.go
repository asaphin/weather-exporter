package main

import (
	"flag"
	"github.com/asaphin/weather-exporter/app"
	"github.com/asaphin/weather-exporter/infrastructure/clients"
	"github.com/asaphin/weather-exporter/infrastructure/http_api"
	"github.com/asaphin/weather-exporter/metrics"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "config", "./weather_exporter.yml", "config file path")

	flag.Parse()

	metrics.Initialize()
}

type Config struct {
	MetricsEndpoint      string               `yaml:"metricsEndpoint"`
	ServerPort           string               `yaml:"serverPort"`
	OpenWeatherMapConfig OpenWeatherMapConfig `yaml:"openWeatherMap"`
}

type OpenWeatherMapConfig struct {
	APIKey   string   `yaml:"apiKey"`
	Location Location `yaml:"location"`
}

type Location struct {
	Lat float64 `yaml:"lat"`
	Lon float64 `yaml:"lon"`
}

func main() {
	log.Info().Str("configFilePath", configFilePath).Msg("config file lookup")

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load config file")
		os.Exit(1)
	}

	log.Info().Msg("config file loaded")

	config := Config{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to parse config file")
		os.Exit(1)
	}

	log.Info().Msg("config unmarshalled")

	openWeatherMapClient := clients.NewOpenWeatherMapClient(config.OpenWeatherMapConfig.APIKey)

	weatherService := app.NewWeatherService(openWeatherMapClient, config.OpenWeatherMapConfig.Location.Lat, config.OpenWeatherMapConfig.Location.Lon)

	httpApi := http_api.NewHttpAPI(config.MetricsEndpoint, config.ServerPort, weatherService)

	log.Info().Msg("running the application...")

	if err = httpApi.Run(); err != nil {
		log.Fatal().Err(err).Msg("server running error")
	}
}
