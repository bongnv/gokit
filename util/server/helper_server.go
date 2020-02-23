package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bongnv/gokit/util/httpserver"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type helperServer struct {
	httpServer    HTTPServer
	httpEndpoints []httpserver.Endpoint
	httpOptions   []httptransport.ServerOption
	middlewares   []endpoint.Middleware
}

func getDefaultNanny() *helperServer {
	return &helperServer{
		httpServer:    httpserver.New(),
		httpEndpoints: nil,
		httpOptions:   nil,
	}
}

func (n *helperServer) getServers() []server {
	servers := []server{}
	if n.httpServer != nil {
		servers = append(servers, n.httpServer)
	}

	return servers
}

func (n *helperServer) initServers() error {
	n.httpServer.WithEndpoint(n.httpEndpoints...)
	n.httpServer.WithOption(n.httpOptions...)

	for _, s := range n.getServers() {
		if err := s.Init(); err != nil {
			return err
		}
	}

	return nil
}

func (n *helperServer) startServers() {
	var wg sync.WaitGroup
	for _, s := range n.getServers() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.Serve(); err != nil {
				// TODO: log errors
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		n.waitForStopSignal()
		n.stopServers()
	}()

	wg.Wait()
}

// waitForStopSignal blocks goroutine until there is a shutdown signal.
func (n *helperServer) waitForStopSignal() {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	<-shutdownCh
	signal.Stop(shutdownCh)
}

func (n *helperServer) stopServers() {
	var wg sync.WaitGroup

	for _, s := range n.getServers() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.Stop(); err != nil {
				// TODO: log errors
			}
		}()
	}

	wg.Wait()
}

func (n *helperServer) run() error {
	if err := n.initServers(); err != nil {
		return err
	}

	n.startServers()
	return nil
}