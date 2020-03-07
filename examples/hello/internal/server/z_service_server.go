package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bongnv/gokit/examples/hello"
	"github.com/bongnv/gokit/examples/hello/internal/endpoint"
	"github.com/bongnv/gokit/util/httpserver"
	"github.com/bongnv/gokit/util/server"
)

// GetServiceOptions ...
func GetServiceOptions(s hello.Service) []server.Option {
	serverEndpoints := endpoint.MakeServiceEndpoints(s)

	opts := []server.Option{}

	opts = append(opts, getServiceHTTPOptions(serverEndpoints)...)

	return opts
}

func getServiceHTTPOptions(serverEndpoints endpoint.ServiceEndpoints) []server.Option {
	opts := []server.Option{}

	opts = append(opts,
		server.WithHTTPEndpoint(httpserver.Endpoint{
			Method:         "",
			Path:           "/hello",
			Endpoint:       serverEndpoints.HelloEndpoint,
			RequestDecoder: decodeServiceHelloRequest,
		}),
		server.WithHTTPEndpoint(httpserver.Endpoint{
			Method:         "",
			Path:           "/bye",
			Endpoint:       serverEndpoints.ByeEndpoint,
			RequestDecoder: decodeServiceByeRequest,
		}),
	)

	return opts
}

func decodeServiceHelloRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req *hello.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, err
}

func decodeServiceByeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req *hello.ByeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, err
}
