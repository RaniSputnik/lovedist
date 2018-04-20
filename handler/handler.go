package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New(buildDir string) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/_ah/ping", pingHandler()).Methods(http.MethodGet)
	router.HandleFunc("/_ah/health", healthHandler()).Methods(http.MethodGet)

	router.HandleFunc("/", indexHandler()).Methods(http.MethodGet)
	router.HandleFunc("/build", buildHandler(buildDir)).Methods(http.MethodPost)

	return wrapGlobalMiddleware(router)
}
