package main

import "net/url"

type IWeatherUrlFactory interface {
	CreateCurrentUrlForCity(city string) (*url.URL, error)
	CreateForecastUrlForCity(city string, days int) (*url.URL, error)
}

type WeatherUrlFactory struct {
	Key string
}

func NewWeatherUrlFactory(key string) *WeatherUrlFactory {
	return &WeatherUrlFactory{
		Key: key,
	}
}

func (w *WeatherUrlFactory) CreateCurrentUrlForCity(city string) (*url.URL, error) {
	params := url.Values{}
	params.Add("key", w.Key)
	params.Add("aqi", "no")
	params.Add("q", city)

	// http://api.weatherapi.com/v1/current.json?key=d3b0abb9a0ba4e2fba5174636231509&q=London&aqi=no

	forcastUrl := &url.URL{
		Scheme:   "http",
		Host:     "api.weatherapi.com",
		Path:     "/v1/current.json",
		RawQuery: params.Encode(),
	}

	return forcastUrl, nil
}

func (w *WeatherUrlFactory) CreateForecastUrlForCity(city string, days int) (*url.URL, error) {
	params := url.Values{}
	params.Add("key", w.Key)
	params.Add("aqi", "no")
	params.Add("q", city)
	params.Add("days", string(rune(days)))

	forcastUrl := &url.URL{
		Scheme:   "http",
		Host:     "api.weatherapi.com",
		Path:     "/v1/forecast.json",
		RawQuery: params.Encode(),
	}

	return forcastUrl, nil
}
