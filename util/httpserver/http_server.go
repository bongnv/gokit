package httpserver

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type Endpoint struct {
	Method          string
	Path            string
	Endpoint        endpoint.Endpoint
	RequestDecoder  httptransport.DecodeRequestFunc
	ResponseEncoder httptransport.EncodeResponseFunc
}

// Server ...
type Server struct {
	endpoints []Endpoint

	httpAddress  string
	httpHandler  http.Handler
	httpListener net.Listener
	options      []httptransport.ServerOption
}

func New() *Server {
	return &Server{}
}

func (s *Server) Init() error {
	s.initializeHandler()

	httpListener, err := net.Listen("tcp", s.httpAddress)
	if err != nil {
		return err
	}

	s.httpListener = httpListener
	return nil
}

func (s *Server) initializeHandler() {
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

	s.httpHandler = r
}

func (s *Server) Serve() error {
	return http.Serve(s.httpListener, s.httpHandler)
}

func (s *Server) Stop() error {
	return s.httpListener.Close()
}

func (s *Server) WithEndpoint(endpoints ...Endpoint) {
	s.endpoints = append(s.endpoints, endpoints...)
}

func (s *Server) WithOption(opts ...httptransport.ServerOption) {
	s.options = append(s.options, opts...)
}

// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// TODO: remove hard-coded content-type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
