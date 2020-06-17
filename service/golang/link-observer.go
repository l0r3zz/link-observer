package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var (
	listeningAddress = flag.String("telemetry.address", "localhost:8003", "Address on which to expose metrics.")
	metricsEndpoint  = flag.String("telemetry.endpoint", "/metrics", "Path under which to expose metrics.")
	url              = [...]string{"https://httpstat.us/200", "https://httpstat.us/503"}
	REQUEST_SUMMARY  = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "link_observer_url_response_ms",
		Help: "Agrigate Response time for URL probe",
	}, []string{"url"})
	RESPONSE_GAUGE = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "link_observer_url_up",
		Help: "Up == 1 (ret 200) Down == 0 (ret 503)",
	}, []string{"url"})
)

func process_request(url string) int {
	timer := prometheus.NewTimer(REQUEST_SUMMARY.WithLabelValues(url))
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	defer timer.ObserveDuration()
	return resp.StatusCode
}

func probeSites() {
	var response_code int
	prometheus.Register(REQUEST_SUMMARY)
	prometheus.Register(RESPONSE_GAUGE)
	go func() {
		for true {
			for _, site := range url {
				time.Sleep(1)
				response_code = process_request(site)
				if response_code == 200 {
					RESPONSE_GAUGE.WithLabelValues(site).Set(1)
				} else {
					RESPONSE_GAUGE.WithLabelValues(site).Set(0)
				}

			}
		}
	}()
}

func main() {
	flag.Parse()
	probeSites()
	http.Handle(*metricsEndpoint, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listeningAddress, nil))

}
