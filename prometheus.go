package main

import "github.com/prometheus/client_golang/prometheus"

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

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	temp.Describe(ch)
	pressure.Describe(ch)
	humidity.Describe(ch)
	wind.Describe(ch)
	clouds.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	res, err := owmResult(apiKey, location)

	temp.WithLabelValues(location).Set(res.Main.Temp)
	temp.Collect(ch)

	pressure.WithLabelValues(location).Set(res.Main.Pressure)
	pressure.Collect(ch)

	humidity.WithLabelValues(location).Set(res.Main.Humidity)
	humidity.Collect(ch)

	wind.WithLabelValues(location).Set(res.Wind.Speed)
	wind.Describe(ch)

	clouds.WithLabelValues(location).Set(res.Clouds.All)
	clouds.Describe(ch)
}
