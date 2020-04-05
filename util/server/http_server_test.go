package server

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_initializeHandler(t *testing.T) {
	expectedBody := "1\n"
	h := &httpServer{
		endpoints: []Endpoint{
			{
				Method: http.MethodGet,
				Path:   "/",
				Endpoint: func(ctx context.Context, request interface{}) (response interface{}, err error) {
					return 1, nil
				},
				RequestDecoder: func(context.Context, *http.Request) (request interface{}, err error) {
					return nil, nil
				},
			},
		},
		handlerDecorators: []handlerDecorator{
			func(h http.Handler) http.Handler {
				fmt.Println("handling")
				return h
			},
		},
	}

	h.initializeHandler()
	s := httptest.NewServer(h.httpHandler)
	defer s.Close()
	resp, err := s.Client().Get(s.URL)
	if err != nil {
		t.Fatalf("unexpected error getting from server: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected a status code of 200, got %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unexpected error reading body: %v", err)
	}
	if !bytes.Equal(body, []byte(expectedBody)) {
		t.Fatalf("response should be hello world, was: %q", string(body))
	}
}
