package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"fmt"

	"strings"

	"github.com/RaniSputnik/lovedist/zip"
)

func main() {
	if len(os.Args) < 3 {
		flag.Usage()
		return
	}

	params := &Params{
		InputDir:  os.Args[1],
		OutputDir: os.Args[2],
		Logger:    log.New(os.Stderr, "", 0),
	}

	if err := Build(params); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Completed successfully")
	}
}

type Params struct {
	InputDir  string
	OutputDir string
	Logger    *log.Logger
}

func Build(params *Params) error {
	if params.Logger == nil {
		params.Logger = doNotLogger()
	}

	if params.InputDir == params.OutputDir {
		return fmt.Errorf("Input directory must != output directory")
	}
	if !strings.HasSuffix(params.InputDir, "/") {
		params.InputDir += "/"
	}
	outfile := filepath.Join(params.OutputDir, fmt.Sprintf("%s.love", filepath.Base(params.InputDir)))
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
	return nil
}

// The default logger used if logger is nil. This saves us having to
// make a nil check everytime we want to log; We set up a logger with
// a writer that does nothing.
func doNotLogger() *log.Logger {
	return log.New(&doNotWriter{}, "", 0)
}

type doNotWriter struct{}

func (*doNotWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
