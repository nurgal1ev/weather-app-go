package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("NO CITY")
var ErrNot200 = errors.New("NOT 200")

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrNoCity
		}
		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{"city": city})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var populationResponse CityPopulationResponse
	json.Unmarshal(body, &populationResponse)
	return !populationResponse.Error
}
