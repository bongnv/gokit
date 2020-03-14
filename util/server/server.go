// Package server provides a helper implementation to help handling multiple services.
package server

type server interface {
	Init() error
	Serve() error
	Stop() error
}

// Option presents an option to enhance server.
type Option func(n *helperServer)

// Serve is the single entry to start serving servers.
func Serve(opts ...Option) error {
	n := getDefaultHelper()

	for _, o := range opts {
		o(n)
	}

	return n.run()
}
