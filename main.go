package main

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	fmt.Println(`
 This is a dummy example of prometheus exporter
  Access: http://127.0.0.1:8081`)

	metricsPath := "/metrics"
	listenAddress := ":8081"

	http.Handle(metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Dummy Exporter</title></head>
             <body>
             <h1>Dummy Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}