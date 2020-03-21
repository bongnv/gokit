package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bongnv/gokit/util/log"
	"github.com/go-kit/kit/endpoint"
)

type helperServer struct {
	httpServer  *httpServer
	middlewares []endpoint.Middleware
}

func getDefaultHelper() *helperServer {
	return &helperServer{
		httpServer: &httpServer{},
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
	for _, s := range n.getServers() {
		if err := s.Init(); err != nil {
			log.Error("message", "Error while initializing service", "error", err)
			return err
		}
	}

	return nil
}

func (n *helperServer) startServers() {
	var wg sync.WaitGroup
	for _, s := range n.getServers() {
		wg.Add(1)
		serviceClone := s
		go func() {
			defer wg.Done()
			if err := serviceClone.Serve(); err != nil {
				log.Error("message", "Error while starting service", "error", err)
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
	for _, s := range n.getServers() {
		if err := s.Stop(); err != nil {
			log.Error("message", "Error while stopping service", "error", err)
		}
	}
}

func (n *helperServer) run() error {
	if err := n.initServers(); err != nil {
		return err
	}

	n.startServers()
	return nil
}
