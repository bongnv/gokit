package server

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_initializeHandler(t *testing.T) {
	expectedBody := "1\n"
	decoraterCalled := 0
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
				decoraterCalled++
				return h
			},
		},
	}

	h.initializeHandler()
	s := httptest.NewServer(h.httpHandler)
	defer s.Close()
	resp, err := s.Client().Get(s.URL)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, []byte(expectedBody), body)
	require.Equal(t, 1, decoraterCalled)
}
