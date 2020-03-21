package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WithHTTPAddress(t *testing.T) {
	opt := WithHTTPAddress("localhost:8080")
	mockServer := &helperServer{}
	opt(mockServer)
	require.Equal(t, "localhost:8080", mockServer.httpServer.httpAddress)
}
