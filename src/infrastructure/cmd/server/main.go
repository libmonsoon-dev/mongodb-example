package main

import (
	"flag"
	"mongodb-example/src/infrastructure"
	"net/http"
	"os"
)

func main() {
	address := flag.String("address", "localhost:8080", "TCP address to listen on")

	flag.Parse()

	handler, err := Init()
	logger := infrastructure.NewStdLogger()

	if err != nil {
		logger.Errorln(err)
		os.Exit(1)
	}

	if err := http.ListenAndServe(*address, handler); err != nil {
		logger.Errorln(err)
		os.Exit(1)
	}
}
