package handlers

import (
	"context"

	"github.com/bongnv/gokit/examples/hello"
)

type handlerImpl struct {
}

func (h *handlerImpl) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return nil, nil
}

func (h *handlerImpl) Bye(ctx context.Context, req *hello.ByeRequest) (*hello.ByeResponse, error) {
	return nil, nil
}

func New() hello.Service {
	return &handlerImpl{}
}
