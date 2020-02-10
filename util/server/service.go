// Package servicenanny provides a helper implementation to help handling multiple services.
package server

import (
	"context"

	"github.com/bongnv/gokit/util/httpserver"
	httptransport "github.com/go-kit/kit/transport/http"
)

type server interface {
	Init(ctx context.Context) error
	Serve(ctx context.Context)
	Stop() error
}

type HTTPServer interface {
	server
	WithEndpoint(endpoints ...httpserver.Endpoint)
	WithOption(opts ...httptransport.ServerOption)
}

type Option func(n *nannyImpl)

func Serve(ctx context.Context, opts ...Option) error {
	n := getDefaultNanny()

	for _, o := range opts {
		o(n)
	}

	return n.Run(ctx)
}
