package httputil

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type sampleRequest struct {
	ID     int    `json:"id"`
	Locale string `json:"locale"`
}

func Test_DecoderRequest(t *testing.T) {
	t.Run("nil-path", func(t *testing.T) {
		require.NotPanics(t, func() {
			err := DecodeRequest(&http.Request{}, nil)
			require.NoError(t, err)
		})
	})

	t.Run("decode-url", func(t *testing.T) {
		mockReq, _ := http.NewRequest(http.MethodGet, "http://example.com/users/100?locale=en", &bytes.Buffer{})
		mockReq = mux.SetURLVars(mockReq, map[string]string{
			"id":     "100",
			"locale": "en",
		})

		req := &sampleRequest{}
		err := DecodeRequest(mockReq, req)
		require.NoError(t, err)
		require.EqualValues(t, 100, req.ID)
		require.EqualValues(t, "en", req.Locale)
	})

	t.Run("decode-body", func(t *testing.T) {
		mockReq, _ := http.NewRequest(http.MethodPost, "http://example.com/users/100", strings.NewReader(`{"locale": "en"}`))
		mockReq = mux.SetURLVars(mockReq, map[string]string{
			"id":     "100",
			"locale": "en",
		})

		req := &sampleRequest{}
		err := DecodeRequest(mockReq, req)
		require.NoError(t, err)
		require.EqualValues(t, 100, req.ID)
		require.EqualValues(t, "en", req.Locale)
	})
}
