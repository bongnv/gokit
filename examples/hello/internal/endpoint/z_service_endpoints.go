package endpoint

import (
	"context"

	"github.com/bongnv/gokit/examples/hello"
	"github.com/go-kit/kit/endpoint"
)

// ServiceEndpoints ...
type ServiceEndpoints struct {
	HelloEndpoint endpoint.Endpoint
	ByeEndpoint   endpoint.Endpoint
}

// MakeServiceEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServiceEndpoints(s hello.Service) ServiceEndpoints {
	return ServiceEndpoints{
		HelloEndpoint: makeServiceHelloEndpoint(s),
		ByeEndpoint:   makeServiceByeEndpoint(s),
	}
}

// makeHelloEndpoint returns an endpoint via the passed service.
func makeServiceHelloEndpoint(s hello.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*hello.Request)
		resp, e := s.Hello(ctx, req)
		return resp, e
	}
}

// makeByeEndpoint returns an endpoint via the passed service.
func makeServiceByeEndpoint(s hello.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*hello.ByeRequest)
		resp, e := s.Bye(ctx, req)
		return resp, e
	}
}
