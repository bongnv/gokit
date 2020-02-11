package main

import (
	"log"

	"github.com/bongnv/gokit/examples/hello/internal/handlers"
	"github.com/bongnv/gokit/examples/hello/internal/server"
)

func main() {
	err := server.Serve(
		handlers.New(),
	)

	log.Println("Service stopped with:", err)
}
