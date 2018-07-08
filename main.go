package main

import (
	"flag"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests")

type Exporter struct{
	gauge prometheus.Gauge
	gaugeVec prometheus.GaugeVec
}

func NewExporter(metricsPrefix string) *Exporter {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:metricsPrefix,
		Name:"gauge_metrics",
		Help:"This is a dummy gauge metric"})

	gaugeVec := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:metricsPrefix,
		Name:"gauge_vec_metric",
		Help:"This is a dummy gauga vece metric"},
		[]string{"myLabel"})

	return &Exporter{
		gauge: gauge,
		gaugeVec: gaugeVec,
	}
}

func (e *Exporter) Collect(ch chan<-prometheus.Metric)  {
	e.gauge.Set(float64(0))
	e.gaugeVec.WithLabelValues("hello").Set(float64(0))
	e.gauge.Collect(ch)
	e.gaugeVec.Collect(ch)
}

func (e *Exporter) Describe(ch chan<-*prometheus.Desc)  {
	e.gauge.Describe(ch)
	e.gaugeVec.Describe(ch)
}

func main() {
	fmt.Println(`
 This is a dummy example of prometheus exporter
  Access: http://127.0.0.1:8081`)

	metricsPrefix := "dummy"
	exporter := NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)


	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))

}