package main

import (
	"log"

	"github.com/bongnv/gokit/examples/hello/internal/handlers"
	"github.com/bongnv/gokit/examples/hello/internal/server"
	gokitServer "github.com/bongnv/gokit/util/server"
)

func main() {
	opts := []gokitServer.Option{}
	opts = append(opts, server.GetServiceOptions(handlers.New())...)

	err := gokitServer.Serve(
		opts...,
	)

	log.Println("Service stopped with:", err)
}
