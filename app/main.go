package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = prometheus.NewCounter(prometheus.CounterOpts{Name: "obsciscope_request_count"})

func main() {
	prometheus.MustRegister(counter)
	logPath := os.Getenv("LOG_PATH")
	f, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter.Inc()
		logger.Println("Request received")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Obsciscope"})
	})
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		for {
			logger.Println("Heartbeat")
			time.Sleep(10 * time.Second)
		}
	}()
	http.ListenAndServe(":8080", nil)
}
