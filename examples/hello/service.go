package hello

import "context"

// Service is a simple interface for a service.
type Service interface {
	Hello(ctx context.Context, p HelloRequest) (HelloResponse, error)
}

// HelloRequest presents a request.
type HelloRequest struct {
	Name string `json:"name"`
}

// HelloResponse ...
type HelloResponse struct {
	Text string `json:"text"`
}
