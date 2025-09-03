package main

import (
	"flag"
	"fmt"
	"weather/geo"
)

func main() {
	city := flag.String("city", "", "City")
	//format := flag.Int("format", 1, "Weather output format")
	flag.Parse()
	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
}
