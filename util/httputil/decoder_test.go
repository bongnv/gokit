package httputil

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DecoderURL(t *testing.T) {
	t.Run("nil-path", func(t *testing.T) {
		require.NotPanics(t, func() {
			DecodeURL(&http.Request{}, nil)
		})
	})

	// implement happy path
}
