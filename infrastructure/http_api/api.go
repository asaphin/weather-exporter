package http_api

import (
	"context"
	"errors"
	"fmt"
	"github.com/asaphin/weather-exporter/app"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type HttpAPI struct {
	server         *http.Server
	weatherService *app.WeatherService
}

func NewHttpAPI(metricsEndpoint, port string, weatherService *app.WeatherService) *HttpAPI {
	mux := http.NewServeMux()

	mux.Handle(metricsEndpoint, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		weatherService.UpdateWeatherData()
		promhttp.Handler().ServeHTTP(w, r)
	}))

	return &HttpAPI{
		server: &http.Server{
			Addr:    port,
			Handler: mux,
		},
	}
}

func (a *HttpAPI) Run() error {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Info().Msg("got interrupt signal, shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("unable to shutdown HTTP API gracefully: %w", err)
	}

	return nil
}
