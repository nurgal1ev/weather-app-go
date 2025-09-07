package weather_test

import (
	"strings"
	"testing"
	"weather/geo"
	"weather/weather"
)

func TestGetWeather(t *testing.T) {
	expected := "Moscow"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("пришла ошибка %v", err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Moscow"
			geoData := geo.GeoData{
				City: expected,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("ожидалось %v, получено %v", weather.ErrWrongFormat, err)
			}
		})
	}
}
