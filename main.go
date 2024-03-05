package main

import (
	"github.com/asaphin/weather-exporter/infrastructure/http_api"
)

func main() {
	httpApi := http_api.HttpAPI{}

	httpApi.Run()
}
