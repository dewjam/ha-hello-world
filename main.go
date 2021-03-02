package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var port string = "8080"
var requestCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "requests_total",
		Help: "The total number of requests received",
	},
	[]string{"code", "path"},
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld)
	r.HandleFunc("/health", healthCheck)
	r.Handle("/metrics", promhttp.Handler())
	r.Use(logHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: r,
	}

	log.Infof("Starting server on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"Protocol":      r.Proto,
			"Path":          r.URL.Path,
			"ContentLength": r.ContentLength,
			"RemoteAddr":    r.RemoteAddr,
			"StatusCode":    "200",
		}).Info("Request Received")

		requestCounter.With(prometheus.Labels{
			"code": "200",
			"path": r.URL.Path,
		}).Inc()
	})
}
