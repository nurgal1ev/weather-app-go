package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	got, err := geo.GetMyLocation(city)

	if err != nil {
		t.Error(err)
	}

	if got.City != expected.City {
		t.Errorf("ожидалось %v, получено %v", expected, got)
	}
}

func TestGetLocationNoCity(t *testing.T) {
	city := "Londonsd"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("ожидалось %v, получили %v", geo.ErrNoCity, err)
	}

}
