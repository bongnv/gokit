package server

import (
	"testing"

	"github.com/rs/cors"
	"github.com/stretchr/testify/require"
)

func Test_WithHTTPAddress(t *testing.T) {
	opt := WithHTTPAddress("localhost:8080")
	mockServer := &helperServer{
		httpServer: &httpServer{},
	}
	opt(mockServer)
	require.Equal(t, "localhost:8080", mockServer.httpServer.httpAddress)
}

func Test_WithCORS(t *testing.T) {
	opt := WithCORS(cors.Options{})
	mockServer := &helperServer{
		httpServer: &httpServer{},
	}
	opt(mockServer)
	require.Len(t, mockServer.httpServer.handlerDecorators, 1)
}
