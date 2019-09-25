package main

import (
	"fmt"
	"log"
	"net/http"
)

// Listening address parts.
var hostname = "0.0.0.0" // or could use e.g. ""
var port = 8080

func main() {
	log.Println("FileServer started...")

	// Serve all files from the root dir down.
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Disable favicon
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./public/favicon.ico")
	// })

	// Start listening
	listenAddress := fmt.Sprintf("%s:%d", hostname, port)
	log.Printf("Listening on address: '%s'\n", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
