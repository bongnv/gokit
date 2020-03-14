package server

// WithHTTPEndpoint ...
func WithHTTPEndpoint(e Endpoint) Option {
	return func(n *helperServer) {
		n.httpServer.endpoints = append(n.httpServer.endpoints, e)
	}
}
