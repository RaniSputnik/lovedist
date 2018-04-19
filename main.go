package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	out := flag.String("out", "./build", "The output directory to write builds to.")
	port := flag.Int("port", 8080, "The port to run the server on.")
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/_ah/ping", pingHandler()).Methods(http.MethodGet)
	router.HandleFunc("/_ah/health", healthHandler()).Methods(http.MethodGet)

	router.HandleFunc("/", indexHandler()).Methods(http.MethodGet)
	router.HandleFunc("/build", buildHandler(*out)).Methods(http.MethodPost)

	portStr := fmt.Sprintf(":%d", *port)
	log.Printf("Now listening at: http://localhost%s", portStr)
	log.Fatal(http.ListenAndServe(portStr, wrapGlobalMiddleware(router)))
}
