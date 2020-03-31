package httputil

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// DecodeURL loads URL vars into the request struct.
func DecodeURL(r *http.Request, req interface{}) error {
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
