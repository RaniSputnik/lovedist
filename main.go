package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"fmt"

	"strings"

	"github.com/RaniSputnik/lovedist/copy"
	"github.com/RaniSputnik/lovedist/zip"
)

func main() {
	pLove := flag.String("love", "/Applications/love.app", "Path to love")
	flag.Parse()

	if len(os.Args) < 3 {
		flag.Usage()
		return
	}

	logger := log.New(os.Stderr, "", 0)
	params := &Params{
		InputDir:   os.Args[1],
		OutputDir:  os.Args[2],
		PathToLove: *pLove,
		Logger:     logger,
	}

	if err := Build(params); err != nil {
		logger.Fatal(err)
	} else {
		logger.Println("Completed successfully")
	}
}

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
	outfile := filepath.Join(params.OutputDir, fmt.Sprintf("%s.love", params.Name))
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
