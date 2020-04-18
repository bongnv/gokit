package httputil

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// decodeURL loads URL vars into the request struct.
func decodeURL(r *http.Request, req interface{}) error {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		return nil
	}

	input := map[string][]string{}
	for k, v := range vars {
		input[k] = []string{v}
	}

	return decoder.Decode(req, input)
}

// DecodeRequest decodes a HTTP to a request object.
func DecodeRequest(r *http.Request, req interface{}) error {
	if req == nil {
		return nil
	}

	decodeURL(r, req)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err != io.EOF {
		return err
	}

	return nil
}
