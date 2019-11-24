package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/RaniSputnik/lovedist/builder/zip"
	"github.com/RaniSputnik/lovedist/dist"
)

func buildHandler(out string, loveDir string, loveVersions []string) http.HandlerFunc {
	if len(loveVersions) == 0 {
		panic("must support one or more love versions")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		file, err := getFile("uploadfile", r)
		if err != nil {
			log.Printf("Failed to get file: %v", err)
			badRequest(w, r, err)
			return
		}
		log.Printf("File = %+v", file)
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

func getFile(name string, r *http.Request) (multipart.File, error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	receivedFiles := r.MultipartForm.File[name]
	if len(receivedFiles) == 0 {
		return nil, fmt.Errorf("no files found")
	}

	loveFile, err := ioutil.TempFile("", "love-")
	if err != nil {
		return nil, err
	}

	if err := zip.ArchiveMultipartFormFiles(receivedFiles, loveFile, nil); err != nil {
		return nil, err
	}

	_, err = loveFile.Seek(0, 0)
	return loveFile, err
}

func doBuild(input io.Reader, out string, loveDir string) (string, error) {
	// TODO generate a proper ID
	id := fmt.Sprintf("%d", time.Now().Unix())

	pathToLoveApp := filepath.Join(loveDir, "osx/love.app")
	output := filepath.Join(buildDir(out, id), "osx")
	_, err := dist.OSX(pathToLoveApp).Build(dist.Project{
		Name:     "TODO",
		BundleID: "com.example.todo",
	}, input, output)
	return id, err

	// params := &builder.Params{
	// 	OutputDir: buildDir(out, id),
	// 	WinParams: &builder.WinParams{
	// 		PathToLoveExe: filepath.Join(loveDir, "win32/love.exe"),
	// 	},
	// 	MacParams: &builder.MacParams{
	// 		PathToLoveApp:    filepath.Join(loveDir, "osx/love.app"),
	// 		BundleIdentifier: "com.example.todo",
	// 	},
	// }
	// err := builder.Build(input, params)
	// return id, err
}

func buildDir(root, id string) string {
	return filepath.Join(root, fmt.Sprintf("build_%s", id))
}
