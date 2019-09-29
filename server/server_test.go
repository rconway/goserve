package server

import (
	"log"
	"testing"
)

func TestServer(t *testing.T) {
	t.Log("testing server instantiation")
	server := NewServer()
	log.Println(server)
}
