package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	// "github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var port string = "8080"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", logHandler(helloWorld))
	r.HandleFunc("/health", healthCheck)
	r.Handle("/metrics", promhttp.Handler())

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
	fmt.Fprintf(w, "OK")
}

func logHandler(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Protocol":      r.Proto,
			"Host":          r.Header["Host"],
			"Path":          r.URL.Path,
			"ContentLength": r.ContentLength,
			"RemoteAddr":    r.RemoteAddr,
		}).Info("Request")
	})
}
