package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New(buildDir string, loveDir string) http.Handler {
	// TODO: Make this configurable
	loveVersions := []string{"11.2.0", "0.10.2"}
	router := mux.NewRouter()

	router.HandleFunc("/_ah/ping", pingHandler()).Methods(http.MethodGet)
	router.HandleFunc("/_ah/health", healthHandler()).Methods(http.MethodGet)
	router.HandleFunc("/_ah/info", infoHandler(loveVersions)).Methods(http.MethodGet)

	router.HandleFunc("/", indexHandler()).Methods(http.MethodGet)
	router.HandleFunc("/build", buildHandler(buildDir, loveDir, loveVersions)).Methods(http.MethodPost)

	return wrapGlobalMiddleware(router)
}
