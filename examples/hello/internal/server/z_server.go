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

// Serve ...
func Serve(s hello.Service, opts ...server.Option) error {
	serverEndpoints := endpoint.MakeServerEndpoints(s)

	opts = append(opts, getHTTPOptions(serverEndpoints)...)

	return server.Serve(opts...)
}

func getHTTPOptions(serverEndpoints endpoint.Endpoints) []server.Option {
	opts := []server.Option{}

	opts = append(opts,
		server.WithHTTPEndpoint(httpserver.Endpoint{
			Method:         "",
			Path:           "/hello",
			Endpoint:       serverEndpoints.HelloEndpoint,
			RequestDecoder: decodeHelloRequest,
		}),
		server.WithHTTPEndpoint(httpserver.Endpoint{
			Method:         "",
			Path:           "/bye",
			Endpoint:       serverEndpoints.ByeEndpoint,
			RequestDecoder: decodeByeRequest,
		}),
	)

	return opts
}

func decodeHelloRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req *hello.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, err
}

func decodeByeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req *hello.ByeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, err
}
