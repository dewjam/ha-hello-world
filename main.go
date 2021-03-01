package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	log.SetFormatter(&log.JSONFormatter{})
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
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
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
		}).Info("Request")

		requestCounter.With(prometheus.Labels{
			"code": "200",
			"path": r.URL.Path,
		}).Inc()
	})
}
