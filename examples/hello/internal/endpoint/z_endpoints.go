package endpoint

import (
	"context"

	"github.com/bongnv/gokit/examples/hello"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	HelloEndpoint endpoint.Endpoint
	ByeEndpoint   endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServerEndpoints(s hello.Service) Endpoints {
	return Endpoints{
		HelloEndpoint: makeHelloEndpoint(s),
		ByeEndpoint:   makeByeEndpoint(s),
	}
}

// makeHelloEndpoint returns an endpoint via the passed service.
func makeHelloEndpoint(s hello.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*hello.HelloRequest)
		resp, e := s.Hello(ctx, req)
		return resp, e
	}
}

// makeByeEndpoint returns an endpoint via the passed service.
func makeByeEndpoint(s hello.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*hello.ByeRequest)
		resp, e := s.Bye(ctx, req)
		return resp, e
	}
}
