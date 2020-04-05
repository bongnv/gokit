package server

import "github.com/rs/cors"

// WithHTTPEndpoint ...
func WithHTTPEndpoint(e Endpoint) Option {
	return func(n *helperServer) {
		n.httpServer.endpoints = append(n.httpServer.endpoints, e)
	}
}

// WithHTTPAddress sets the address for HTTP service to listen on.
func WithHTTPAddress(addr string) Option {
	return func(n *helperServer) {
		n.httpServer.httpAddress = addr
	}
}

// WithCORS adds CORS support.
func WithCORS(opts cors.Options) Option {
	return func(n *helperServer) {
		c := cors.New(opts)
		n.httpServer.handlerDecorators = append(
			n.httpServer.handlerDecorators,
			c.Handler,
		)
	}
}
