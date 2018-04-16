package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RaniSputnik/lovedist/builder"
)

func buildHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		if err := doBuild(file); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO write HTML
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func doBuild(input io.Reader) error {
	logger := log.New(os.Stderr, "", 0)
	params := &builder.Params{
		OutputDir: fmt.Sprintf("./tmp/build_%d", time.Now().Unix()),
		Logger:    logger,
		WinParams: &builder.WinParams{
			PathToLoveExe: "./love/win32/love.exe",
		},
	}
	return builder.Build(input, params)
}
