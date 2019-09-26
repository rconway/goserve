package server

import (
	"net/http"
)

// GoServer implements a simple file server
type GoServer struct {
	http.ServeMux
}

// NewServer creates a GoServer
func NewServer() *GoServer {
	server := GoServer{}

	// Serve all files from the root dir down.
	server.Handle("/", http.FileServer(http.Dir(".")))

	// Disable favicon
	server.Handle("/favicon.ico", http.NotFoundHandler())
	// http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./public/favicon.ico")
	// })

	return &server
}

// ListenAndServe starts listening on the supplied address
func (server *GoServer) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, server)
}
