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
	input := os.Args[1]
	output := os.Args[2]
	log.Printf(input)
	if err := doTheWork(input, output); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Completed successfully")
	}
}

func doTheWork(input, output string) error {
	if input == output {
		return fmt.Errorf("Input must != output")
	}
	if !strings.HasSuffix(input, "/") {
		input += "/"
	}
	outfile := filepath.Join(output, fmt.Sprintf("%s.love", filepath.Base(input)))
	log.Printf("Outputting to %s", outfile)
	fw, err := os.Create(outfile)
	if err != nil {
		return err
	}
	if err := zip.Archive(input, fw, logArchivePath); err != nil {
		return err
	}
	return nil
}

func logArchivePath(archivePath string) {
	log.Println(archivePath)
}
