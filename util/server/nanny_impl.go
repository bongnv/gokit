package server

import (
	"context"
	"sync"

	"github.com/bongnv/gokit/util/httpserver"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type nannyImpl struct {
	httpServer    HTTPServer
	httpEndpoints []httpserver.Endpoint
	httpOptions   []httptransport.ServerOption
	middlewares   []endpoint.Middleware
}

func getDefaultNanny() *nannyImpl {
	return &nannyImpl{
		httpServer:    httpserver.New(),
		httpEndpoints: nil,
		httpOptions:   nil,
	}
}

func (n *nannyImpl) getServers() []server {
	servers := []server{}
	if n.httpServer != nil {
		servers = append(servers, n.httpServer)
	}

	return servers
}

func (n *nannyImpl) Init(ctx context.Context) error {
	n.httpServer.WithEndpoint(n.httpEndpoints...)
	n.httpServer.WithOption(n.httpOptions...)

	for _, s := range n.getServers() {
		if err := s.Init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (n *nannyImpl) Serve(ctx context.Context) {
	var wg sync.WaitGroup
	for _, s := range n.getServers() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Serve(ctx)
		}()
	}
	wg.Wait()
}

func (n *nannyImpl) Stop() error {
	for _, s := range n.getServers() {
		if err := s.Stop(); err != nil {
			return err
		}
	}

	return nil
}

func (n *nannyImpl) Run(ctx context.Context) error {
	if err := n.Init(ctx); err != nil {
		return err
	}

	n.Serve(ctx)
	return nil
}
