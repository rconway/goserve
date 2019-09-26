package main

import (
	"fmt"
	"log"

	goserver "github.com/rconway/goserve/server"
)

// Listening address parts.
var hostname = "0.0.0.0" // or could use e.g. ""
var port = 8080

func main() {
	log.Println("FileServer started...")

	server := goserver.NewServer()

	// Start listening
	listenAddress := fmt.Sprintf("%s:%d", hostname, port)
	log.Printf("Listening on address: '%s'\n", listenAddress)
	log.Fatal(server.ListenAndServe(listenAddress))
}
