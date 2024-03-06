# weather-exporter
Weather exporter for Prometheus testing purposes

Build it:
```shell
docker build . --tag=weather-exporter
```

Run it:
```shell
docker run -d --network host -v ./weather_exporter.yml:/etc/weather_exporter/weather_exporter.yml weather-exporter
```
