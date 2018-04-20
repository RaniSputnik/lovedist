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
	love := flag.String("love", "./love", "Path to the Love executables.")
	port := flag.Int("port", 8080, "The port to run the server on.")
	flag.Parse()

	resolvedOut := mustResolve(*out)
	resolvedLove := mustResolve(*love)
	portStr := fmt.Sprintf(":%d", *port)

	log.Printf("Now listening at: http://localhost%s", portStr)
	log.Fatal(http.ListenAndServe(portStr, handler.New(resolvedOut, resolvedLove)))
}

func mustResolve(path string) string {
	result, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return result
}
