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

func buildHandler(out string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		id, err := doBuild(file, out)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			// TODO write HTML
			return
		}

		outDir := buildDir(out, id)
		filename := filepath.Base(outDir)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", filename))
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

		zip.Archive(outDir, w, nil)
	}
}

func doBuild(input io.Reader, out string) (string, error) {
	// TODO generate a proper ID
	id := fmt.Sprintf("%d", time.Now().Unix())

	params := &builder.Params{
		OutputDir: buildDir(out, id),
		WinParams: &builder.WinParams{
			PathToLoveExe: "./love/win32/love.exe",
		},
		MacParams: &builder.MacParams{
			PathToLoveApp:    "./love/osx/love.app",
			BundleIdentifier: "com.example.todo",
		},
	}
	err := builder.Build(input, params)
	return id, err
}

func buildDir(root, id string) string {
	return filepath.Join(root, fmt.Sprintf("build_%s", id))
}
