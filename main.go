package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	listenAddress = flag.String("web.listen", ":9120", "Address on which to expose metrics and web interface.")
	apiKey        = flag.String("apikey", "", "OpenWeatherMap API key")
	location      = flag.String("location", "Dresden,DE", "Geografical location for requested data")
)

func main() {
	flag.Parse()

	prometheus.MustRegister(&Exporter{Location: *location, APIKey: *apiKey})
	log.Printf("Starting Server: %s", *listenAddress)
	handler := prometheus.Handler()
	http.Handle("/metrics", handler)

	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
