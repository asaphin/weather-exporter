package metrics

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	temperature   *prometheus.GaugeVec
	pressure      *prometheus.GaugeVec
	humidity      *prometheus.GaugeVec
	windSpeed     *prometheus.GaugeVec
	windDirection *prometheus.GaugeVec
	cloudsAll     *prometheus.GaugeVec
	rainOneHour   *prometheus.GaugeVec
}

var m = &metrics{
	temperature: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "temperature",
		Help: "Current temperature in Kelvins",
	},
		[]string{"location"}),
	pressure: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pressure",
		Help: "Current pressure in Hectopascals",
	},
		[]string{"location"}),
	humidity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "humidity",
		Help: "Current relative humidity in Percents",
	},
		[]string{"location"}),
	windSpeed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wind_speed",
		Help: "Current wind speed in Meters per Second",
	},
		[]string{"location"}),
	windDirection: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wind_direction",
		Help: "Current wind direction in Degrees (Azimuth)",
	},
		[]string{"location"}),
	cloudsAll: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "clouds_all",
		Help: "Cloudiness in Percents",
	},
		[]string{"location"}),
	rainOneHour: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "rain_1h",
		Help: "Rain volume for the last 1 hour in mm",
	},
		[]string{"location"}),
}

func Initialize() {
	prometheus.MustRegister(
		m.temperature,
		m.pressure,
		m.humidity,
		m.windSpeed,
		m.windDirection,
		m.cloudsAll,
		m.rainOneHour,
	)
}

func RegisterTemperature(temperature float64, locationName string) {
	m.temperature.WithLabelValues(locationName).Set(temperature)
}

func RegisterPressure(pressure float64, locationName string) {
	m.pressure.WithLabelValues(locationName).Set(pressure)
}

func RegisterHumidity(humidity float64, locationName string) {
	m.humidity.WithLabelValues(locationName).Set(humidity)
}

func RegisterWindSpeed(windSpeed float64, locationName string) {
	m.windSpeed.WithLabelValues(locationName).Set(windSpeed)
}

func RegisterWindDirection(windDirection float64, locationName string) {
	m.windDirection.WithLabelValues(locationName).Set(windDirection)
}

func RegisterCloudsAll(cloudsAll float64, locationName string) {
	m.cloudsAll.WithLabelValues(locationName).Set(cloudsAll)
}

func RegisterRainOneHour(rainOneHour float64, locationName string) {
	m.rainOneHour.WithLabelValues(locationName).Set(rainOneHour)
}
