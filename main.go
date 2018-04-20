package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/RaniSputnik/lovedist/handler"
)

func main() {
	out := flag.String("out", "./build", "The output directory to write builds to.")
	port := flag.Int("port", 8080, "The port to run the server on.")
	flag.Parse()

	resolvedOut, err := filepath.Abs(*out)
	if err != nil {
		log.Fatalf("Failed to resolve 'out' parameter: %s", err)
	}
	portStr := fmt.Sprintf(":%d", *port)

	log.Printf("Now listening at: http://localhost%s", portStr)
	log.Fatal(http.ListenAndServe(portStr, handler.New(resolvedOut)))
}
