package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/RaniSputnik/lovedist/builder"
	"github.com/RaniSputnik/lovedist/builder/zip"
)

func buildHandler(out string, loveDir string, loveVersions []string) http.HandlerFunc {
	if len(loveVersions) == 0 {
		panic("must support one or more love versions")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		file, err := getFile("uploadfile", r)
		if err != nil {
			badRequest(w, r, err)
			return
		}
		defer file.Close()

		// TODO: Instead read this from conf.lua if it exists
		// and if it doesn't exist, just use the latest version
		loveVersion := r.FormValue("loveversion")
		if loveVersion == "" {
			loveVersion = loveVersions[0]
		}
		if err := loveVersionOK(loveVersion, loveVersions); err != nil {
			badRequest(w, r, err)
			return
		}

		id, err := doBuild(file, out, filepath.Join(loveDir, loveVersion))
		if err != nil {
			internalServerError(w, r, err)
			return
		}

		outDir := buildDir(out, id)
		filename := filepath.Base(outDir)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", filename))
		// TODO: Sanitize input Content-Type here
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

		zip.Archive(outDir, w, nil)
	}
}

func loveVersionOK(version string, supportedVersions []string) error {
	for _, v := range supportedVersions {
		if v == version {
			return nil
		}
	}
	return fmt.Errorf("Unsupported love version: '%s', supported versions are '%s'",
		version, supportedVersions)
}

func getFile(name string, r *http.Request) (file multipart.File, err error) {
	if err = r.ParseMultipartForm(32 << 20); err != nil {
		return file, err
	}
	file, _, err = r.FormFile(name)
	return file, err
}

func doBuild(input io.Reader, out string, loveDir string) (string, error) {
	// TODO generate a proper ID
	id := fmt.Sprintf("%d", time.Now().Unix())

	params := &builder.Params{
		OutputDir: buildDir(out, id),
		WinParams: &builder.WinParams{
			PathToLoveExe: filepath.Join(loveDir, "win32/love.exe"),
		},
		MacParams: &builder.MacParams{
			PathToLoveApp:    filepath.Join(loveDir, "osx/love.app"),
			BundleIdentifier: "com.example.todo",
		},
	}
	err := builder.Build(input, params)
	return id, err
}

func buildDir(root, id string) string {
	return filepath.Join(root, fmt.Sprintf("build_%s", id))
}
