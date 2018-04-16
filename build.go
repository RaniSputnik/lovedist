package main

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/RaniSputnik/lovedist/builder"
	"github.com/RaniSputnik/lovedist/builder/zip"
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

		id, err := doBuild(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO write HTML
			return
		}

		filename := filepath.Base(buildDir(id))
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", filename))
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

		zip.Archive(buildDir(id), w, nil)
	}
}

func doBuild(input io.Reader) (string, error) {
	// TODO generate a proper ID
	id := fmt.Sprintf("%d", time.Now().Unix())

	params := &builder.Params{
		OutputDir: buildDir(id),
		WinParams: &builder.WinParams{
			PathToLoveExe: "./love/win32/love.exe",
		},
	}
	err := builder.Build(input, params)
	return id, err
}

func buildDir(id string) string {
	return fmt.Sprintf("./tmp/build_%s", id)
}
