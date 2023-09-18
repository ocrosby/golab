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

func GetCurrentWeather(city string, factory WeatherUrlFactory) (*Weather, error) {
	atlantaUrl, err := factory.CreateCurrentUrlForCity(city)
	if err != nil {
		return nil, err
	}

	context := context2.Background()
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(context, http.MethodGet, atlantaUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Request: %s %s\n", req.Method, req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// convert byte slice to io.Reader
	reader := bytes.NewReader(body)

	fmt.Printf("Status Code: %d\n", res.StatusCode)
	fmt.Printf("Status: %s\n", res.Status)

	if res.StatusCode != http.StatusOK {
		var weatherError WeatherError

		err = json.NewDecoder(reader).Decode(&weatherError)
		if err != nil {
			return nil, err
		}

		fmt.Printf("weather API returned an error: %d %s\n", weatherError.Error.Code, weatherError.Error.Message)
		panic("weather API not available")
	}

	var weather Weather

	if err = json.NewDecoder(reader).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func GetForcast(city string, days int, factory WeatherUrlFactory) (*Weather, error) {
	atlantaUrl, err := factory.CreateForecastUrlForCity(city, days)
	if err != nil {
		return nil, err
	}

	context := context2.Background()
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(context, http.MethodGet, atlantaUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Request: %s %s\n", req.Method, req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// convert byte slice to io.Reader
	reader := bytes.NewReader(body)

	fmt.Printf("Status Code: %d\n", res.StatusCode)
	fmt.Printf("Status: %s\n", res.Status)

	if res.StatusCode != http.StatusOK {
		var weatherError WeatherError

		err = json.NewDecoder(reader).Decode(&weatherError)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("weather API returned an error: %d %s", weatherError.Error.Code, weatherError.Error.Message)
	}

	var weather Weather

	if err = json.NewDecoder(reader).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func DisplayWeatherForCity(city string, factory *WeatherUrlFactory) error {
	weather, err := GetCurrentWeather(city, *factory)
	if err != nil {
		return err
	}

	fmt.Printf("Weather for %s\n", weather.Location.Name)

	return nil
}

func DisplayForcastForCity(city string, days int, factory *WeatherUrlFactory) error {
	weather, err := GetForcast(city, days, *factory)
	if err != nil {
		return err
	}

	for _, day := range weather.Forecast.Forecastday {
		for _, hour := range day.Hour {
			fmt.Printf("%s: %s %.2f %.2f\n", hour.TimeEpoch, hour.Condition.Text, hour.TempC, hour.TempF)
		}
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

	err := DisplayForcastForCity("Los Angeles", 3, factory)
	if err != nil {
		panic(err)
	}
}
