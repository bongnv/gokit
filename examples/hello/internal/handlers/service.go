package handlers

import (
	"context"

	"github.com/bongnv/gokit/examples/hello"
)

type handlerImpl struct {
}

func (h *handlerImpl) Hello(ctx context.Context, req *hello.Request) (*hello.Response, error) {
	return nil, nil
}

func (h *handlerImpl) Bye(ctx context.Context, req *hello.ByeRequest) (*hello.ByeResponse, error) {
	return nil, nil
}

// New creates a new service with business logic.
func New() hello.Service {
	return &handlerImpl{}
}
