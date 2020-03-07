// Package server provides a helper implementation to help handling multiple services.
package server

import (
	"github.com/bongnv/gokit/util/httpserver"
	httptransport "github.com/go-kit/kit/transport/http"
)

type server interface {
	Init() error
	Serve() error
	Stop() error
}

// HTTPServer requires method WithEndpoint and WithOption on top of a service.
type HTTPServer interface {
	server
	WithEndpoint(endpoints ...httpserver.Endpoint)
	WithOption(opts ...httptransport.ServerOption)
}

// Option presents an option to enhance server.
type Option func(n *helperServer)

// Serve is the single entry to start serving servers.
func Serve(opts ...Option) error {
	n := getDefaultNanny()

	for _, o := range opts {
		o(n)
	}

	return n.run()
}
