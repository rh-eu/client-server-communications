package main

import (
	"log"

	"github.com/rh-eu/client-server-communications/grpc/go/route-guide-example/pkg/server"
)

func main() {
	server := server.NewServer()

	log.Printf("Starting server ... %+v", server)

	server.Run()
}
