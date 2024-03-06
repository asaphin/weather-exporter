package clients

import (
	"encoding/json"
	"fmt"
	"github.com/asaphin/weather-exporter/domain"
	"io"
	"net/http"
)

type CurrentWeatherDTO struct {
	Coordinates Coordinates        `json:"coord"`
	Weather     []WeatherItem      `json:"weather"`
	Base        string             `json:"base"`
	Main        MainWeather        `json:"main"`
	Visibility  float64            `json:"visibility"`
	Wind        Wind               `json:"wind"`
	Rain        Rain               `json:"rain"`
	Clouds      Clouds             `json:"clouds"`
	Dt          int                `json:"dt"`
	System      LocationSystemData `json:"sys"`
	Timezone    int                `json:"timezone"`
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Cod         int                `json:"cod"`
}

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type WeatherItem struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainWeather struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type Rain struct {
	OneHour float64 `json:"1h"`
}

type Clouds struct {
	All float64 `json:"all"`
}

type LocationSystemData struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type OpenWeatherMapClient struct {
	apiKey string
	client *http.Client
}

func (c *OpenWeatherMapClient) GetWeather(lat, lon float64) (*domain.Weather, error) {
	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.2f&lon=%.2f&appid=%s", lat, lon, c.apiKey)

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("error making the request:", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading the response body:", err)
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("response status %d: %s", response.StatusCode, string(body))
	}

	weather := CurrentWeatherDTO{}

	err = json.Unmarshal(body, &weather)

	return &domain.Weather{
		LocationName:  weather.Name,
		Latitude:      weather.Coordinates.Latitude,
		Longitude:     weather.Coordinates.Longitude,
		Temperature:   weather.Main.Temp,
		Pressure:      weather.Main.Pressure,
		Humidity:      weather.Main.Humidity,
		WindSpeed:     weather.Wind.Speed,
		WindDirection: weather.Wind.Deg,
		RainOneHour:   weather.Rain.OneHour,
		CloudsAll:     weather.Clouds.All,
	}, nil
}

func NewOpenWeatherMapClient(apiKey string) *OpenWeatherMapClient {
	return &OpenWeatherMapClient{
		apiKey: apiKey,
		client: http.DefaultClient,
	}
}
