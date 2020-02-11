package server

import "github.com/bongnv/gokit/util/httpserver"

func WithHTTPEndpoint(e httpserver.Endpoint) Option {
	return func(n *helperServer) {
		n.httpEndpoints = append(n.httpEndpoints, e)
	}
}

func WithHTTPServer(s HTTPServer) Option {
	return func(n *helperServer) {
		if s != nil {
			n.httpServer = s
		}
	}
}
