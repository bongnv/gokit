package generator

import (
	"bytes"
	"testing"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/stretchr/testify/require"
)

type bufWriter struct {
	b bytes.Buffer
}

func (w *bufWriter) Write(path string, data []byte) error {
	_, _ = w.b.Write([]byte(path + ":\n"))
	_, err := w.b.Write(data)
	return err
}

func Test_fileGenerator(t *testing.T) {
	mockWriter := &bufWriter{}
	generator := &Generator{
		FilePath:     "main.go",
		TemplateName: "main",
		Service: &parser.Service{
			Package:     "github.com/hello",
			PackageName: "hello",
		},
		Writer:                 mockWriter,
		WithAutogeneratedNotes: true,
	}

	err := generator.Do()
	require.NoError(t, err)
	require.Equal(t,
		`main.go:
// Code generated by gokit v0.0.1. DO NOT EDIT.

package main

import (
	"log"

	gokitServer "github.com/bongnv/gokit/util/server"
	"github.com/hello/internal/handlers"
	"github.com/hello/internal/service"
)

func main() {
	opts := []gokitServer.Option{}
	opts = append(opts, service.GetOptions(handlers.New())...)

	err := gokitServer.Serve(
		opts...,
	)

	log.Println("Service stopped with:", err)
}
`,
		mockWriter.b.String(),
	)
}