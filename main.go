package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "Listening port")
	flag.Parse()
}

func main() {
	logger := RequestLogger{
		fieldColor:  colorGreen,
		valueColour: colorYellow,
	}

	http.HandleFunc("/", handler(logger))

	address := fmt.Sprintf(":%d", port)
	log.Println("Listen on", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v\n", err)
	}
}
