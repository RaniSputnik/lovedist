package main

import (
	"flag"
	"log"
	"os"
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
