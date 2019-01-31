package handler

import (
	"encoding/json"
	"net/http"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

func infoHandler(loveVersions []string) http.HandlerFunc {
	type Info struct {
		Love struct {
			SupportedVersions []string `json:"supported_versions"`
		} `json:"love"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var info Info
		info.Love.SupportedVersions = loveVersions
		json.NewEncoder(w).Encode(info)
	}
}
