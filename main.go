package main

import (
	"flag"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"fmt"
)

var addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests")

func main() {
	fmt.Println(`
 This is a dummy example of prometheus exporter
  Access: http://127.0.0.1:8081`)

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))

}