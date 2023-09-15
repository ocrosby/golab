/*
This program reads weather information from the weather API.

https://www.weatherapi.com/
*/
package main

import (
	"bytes"
	context2 "context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func GetCurrentWeather(city string, factory WeatherUrlFactory) error {
	atlantaUrl, err := factory.CreateCurrentUrlForCity(city)
	if err != nil {
		return err
	}

	context := context2.Background()
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(context, http.MethodGet, atlantaUrl.String(), nil)
	if err != nil {
		return err
	}

	fmt.Printf("Request: %s %s\n", req.Method, req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// convert byte slice to io.Reader
	reader := bytes.NewReader(body)

	fmt.Printf("Status Code: %d\n", res.StatusCode)
	fmt.Printf("Status: %s\n", res.Status)

	if res.StatusCode != http.StatusOK {
		var weatherError WeatherError

		err = json.NewDecoder(reader).Decode(&weatherError)
		if err != nil {
			return err
		}

		fmt.Printf("weather API returned an error: %d %s\n", weatherError.Error.Code, weatherError.Error.Message)
		panic("weather API not available")
	}

	var weather Weather

	err = json.NewDecoder(reader).Decode(&weather)
	if err != nil {
		return err
	}

	fmt.Println(weather)

	return nil
}

func GetForcast(city string, days int, factory WeatherUrlFactory) error {
	atlantaUrl, err := factory.CreateForecastUrlForCity(city, days)
	if err != nil {
		return err
	}

	context := context2.Background()
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(context, http.MethodGet, atlantaUrl.String(), nil)
	if err != nil {
		return err
	}

	fmt.Printf("Request: %s %s\n", req.Method, req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// convert byte slice to io.Reader
	reader := bytes.NewReader(body)

	fmt.Printf("Status Code: %d\n", res.StatusCode)
	fmt.Printf("Status: %s\n", res.Status)

	if res.StatusCode != http.StatusOK {
		var weatherError WeatherError

		err = json.NewDecoder(reader).Decode(&weatherError)
		if err != nil {
			return err
		}

		fmt.Printf("weather API returned an error: %d %s\n", weatherError.Error.Code, weatherError.Error.Message)
		panic("weather API not available")
	}

	var weatherSlice []Weather

	err = json.NewDecoder(reader).Decode(&weatherSlice)
	if err != nil {
		return err
	}

	for _, weather := range weatherSlice {
		fmt.Println(weather)
	}

	return nil
}

func DisplayWeatherForCity(city string, factory *WeatherUrlFactory) error {
	err := GetCurrentWeather(city, *factory)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	key := os.Getenv("WEATHER_API_KEY")
	factory := NewWeatherUrlFactory(key)

	_ = DisplayWeatherForCity("London", factory)
	_ = DisplayWeatherForCity("Atlanta", factory)
	_ = DisplayWeatherForCity("Raleigh", factory)
	_ = DisplayWeatherForCity("San Jose", factory)
}
