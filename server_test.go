package main

import (
	"log"
	"testing"

	"github.com/rconway/goserve/server"
)

func TestServer(t *testing.T) {
	t.Log("testing server instantiation")
	server := server.NewServer()
	log.Println(server)
}
