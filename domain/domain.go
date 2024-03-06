package domain

type Weather struct {
	LocationName  string
	Latitude      float64
	Longitude     float64
	Temperature   float64
	Pressure      float64
	Humidity      float64
	WindSpeed     float64
	WindDirection float64
	CloudsAll     float64
	RainOneHour   float64
}
