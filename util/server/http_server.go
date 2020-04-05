package server

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/bongnv/gokit/util/log"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Endpoint presents an endpoint.
type Endpoint struct {
	Method          string
	Path            string
	Endpoint        endpoint.Endpoint
	RequestDecoder  httptransport.DecodeRequestFunc
	ResponseEncoder httptransport.EncodeResponseFunc
}

// Init ...
func (s *httpServer) Init() error {
	s.initializeHandler()

	httpListener, err := net.Listen("tcp", s.httpAddress)
	if err != nil {
		return err
	}

	s.httpListener = httpListener
	return nil
}

// Serve ...
func (s *httpServer) Serve() error {
	log.Info("msg", "HTTP service is starting at "+s.httpListener.Addr().String())
	return http.Serve(s.httpListener, s.httpHandler)
}

// Stop ...
func (s *httpServer) Stop() error {
	return s.httpListener.Close()
}

// WithEndpoint ...
func (s *httpServer) WithEndpoint(endpoints ...Endpoint) {
	s.endpoints = append(s.endpoints, endpoints...)
}

// WithOption ...
func (s *httpServer) WithOption(opts ...httptransport.ServerOption) {
	s.options = append(s.options, opts...)
}

type handlerDecorator func(h http.Handler) http.Handler

type httpServer struct {
	endpoints         []Endpoint
	options           []httptransport.ServerOption
	httpAddress       string
	handlerDecorators []handlerDecorator

	httpHandler  http.Handler
	httpListener net.Listener
}

func (s *httpServer) initializeHandler() {
	r := mux.NewRouter()

	for _, e := range s.endpoints {
		route := r.NewRoute()
		if e.Method != "" {
			route.Methods(e.Method)
		}

		encoder := encodeHTTPGenericResponse
		if e.ResponseEncoder != nil {
			encoder = e.ResponseEncoder
		}

		route.Path(e.Path).Handler(httptransport.NewServer(
			e.Endpoint,
			e.RequestDecoder,
			encoder,
			s.options...,
		))
	}

	for _, d := range s.handlerDecorators {
		s.httpHandler = d(r)
	}
}

// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// TODO: remove hard-coded content-type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
