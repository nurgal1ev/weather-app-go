package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {
	city := flag.String("city", "", "City")
	format := flag.Int("format", 1, "Weather output format")

	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData.City)

	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
