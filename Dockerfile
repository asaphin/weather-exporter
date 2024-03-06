FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o weather_exporter


FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

COPY --from=builder /app/weather_exporter .
COPY --from=builder /app/weather_exporter.yml /etc/weather_exporter/weather_exporter.yml

EXPOSE 9110

CMD ["./weather_exporter", "-config", "/etc/weather_exporter/weather_exporter.yml"]
