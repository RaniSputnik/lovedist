package handler

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func wrapGlobalMiddleware(h http.Handler) http.Handler {
	h = handlers.LoggingHandler(os.Stdout, h)
	h = handlers.RecoveryHandler()(h)
	return h
}
