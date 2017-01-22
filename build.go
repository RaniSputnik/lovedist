package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/RaniSputnik/lovedist/copy"
	"github.com/pierrre/archivefile/zip"
)

type Params struct {
	Name       string
	InputDir   string
	OutputDir  string
	PathToLove string
	Logger     *log.Logger
}

func Build(params *Params) error {

	// Validate params
	if params.Logger == nil {
		params.Logger = doNotLogger()
	}
	if params.InputDir == params.OutputDir {
		return fmt.Errorf("Input directory must != output directory")
	}
	if !strings.HasSuffix(params.InputDir, "/") {
		params.InputDir += "/"
	}
	if params.Name == "" {
		params.Name = filepath.Base(params.InputDir)
	}

	// Copy the love.app
	outapp := filepath.Join(params.OutputDir, fmt.Sprintf("%s.app", params.Name))
	if err := copy.Dir(params.PathToLove, outapp); err != nil {
		return err
	}

	// Create the .love file
	outfilename := fmt.Sprintf("%s.love", params.Name)
	outfile := filepath.Join(params.OutputDir, outfilename)
	params.Logger.Printf("Outputting to %s", outfile)
	fw, err := os.Create(outfile)
	if err != nil {
		return err
	}
	err = zip.Archive(params.InputDir, fw, func(archivePath string) {
		params.Logger.Printf("Zipping %s", archivePath)
	})
	if err != nil {
		return err
	}

	// Copy .love file into love app
	// TODO we have kept this a separate step because we could
	// perform "Copy love.app" and "Create .love" steps concurrently
	finallovepath := filepath.Join(outapp, "Contents", "Resources", outfilename)
	if err := copy.File(outfile, finallovepath); err != nil {
		return err
	}

	return nil
}
