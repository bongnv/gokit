package server

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
