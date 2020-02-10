package main

import (
	"context"
	"log"

	"github.com/bongnv/gokit/examples/hello/internal/handlers"
	"github.com/bongnv/gokit/examples/hello/internal/server"
)

func main() {
	err := server.Serve(context.Background(),
		handlers.New(),
	)

	if err != nil {
		log.Println(err)
	}
}
