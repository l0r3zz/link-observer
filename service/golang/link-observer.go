// NOTE: This is not yet functional!
package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	//	"time"
)

var (
	url             = [...]string{"https://httpstat.us/200", "https://httpstat.us/503"}
	REQUEST_SUMMARY = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "link_observer_url_response_ms",
		Help: "Agrigate Response time for URL probe",
	}, []string{"url"})
	RESPONSE_GAUGE = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "link_observer_url_up",
		Help: "Up == 1 (ret 200) Down == 0 (ret 503)",
	}, []string{"url"})
)

func process_request(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get failed to %s\n", url)
		return -1
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8003", nil)
	var response_code int
	fmt.Printf("Got  to %s\n", url)
	for _, site := range url {
		//		time.Sleep(1)
		response_code = process_request(site)
		if response_code == 200 {
			fmt.Printf("OK")
		} else {
			fmt.Printf("BAD")
		}

	}
}
