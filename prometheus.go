package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type result struct {
	Main struct {
		Temp     float64
		Pressure float64
		Humidity float64
	}
	Wind struct {
		Speed float64
	}
	Clouds struct {
		All float64
	}
}

var (
	temp = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "openweathermap",
		Name:      "temperature_celsius",
		Help:      "Temperature in °C",
	}, []string{"location"})

	pressure = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "openweathermap",
		Name:      "pressure_hpa",
		Help:      "Atmospheric pressure in hPa",
	}, []string{"location"})

	humidity = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "openweathermap",
		Name:      "humidity_percent",
		Help:      "Humidity in Percent",
	}, []string{"location"})

	wind = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "openweathermap",
		Name:      "wind_mps",
		Help:      "Wind speed in m/s",
	}, []string{"location"})

	clouds = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "openweathermap",
		Name:      "cloudiness_percent",
		Help:      "Cloudiness in Percent",
	}, []string{"location"})
)

type Exporter struct {
	Location string
	APIKey   string
}

func (e *Exporter) owmData() (result, error) {
	var res result
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", e.Location, e.APIKey)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&res)
	return res, err
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	temp.Describe(ch)
	pressure.Describe(ch)
	humidity.Describe(ch)
	wind.Describe(ch)
	clouds.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	res, err := e.owmData()
	if err != nil {
		log.Printf("could not get results from OpenWeatherMap: %v", err)
		return
	}

	temp.WithLabelValues(e.Location).Set(res.Main.Temp)
	temp.Collect(ch)

	pressure.WithLabelValues(e.Location).Set(res.Main.Pressure)
	pressure.Collect(ch)

	humidity.WithLabelValues(e.Location).Set(res.Main.Humidity)
	humidity.Collect(ch)

	wind.WithLabelValues(e.Location).Set(res.Wind.Speed)
	wind.Collect(ch)

	clouds.WithLabelValues(e.Location).Set(res.Clouds.All)
	clouds.Collect(ch)
}
