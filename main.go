package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/_ah/ping", pingHandler()).Methods(http.MethodGet)
	router.HandleFunc("/_ah/health", healthHandler()).Methods(http.MethodGet)

	router.HandleFunc("/", indexHandler()).Methods(http.MethodGet)
	router.HandleFunc("/build", buildHandler()).Methods(http.MethodPost)

	http.ListenAndServe(port, wrapGlobalMiddleware(router))
}
