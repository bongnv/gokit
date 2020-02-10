package server

import "github.com/bongnv/gokit/util/httpserver"

func WithHTTPEndpoint(e httpserver.Endpoint) Option {
	return func(n *nannyImpl) {
		n.httpEndpoints = append(n.httpEndpoints, e)
	}
}

func WithHTTPServer(s HTTPServer) Option {
	return func(n *nannyImpl) {
		if s != nil {
			n.httpServer = s
		}
	}
}
