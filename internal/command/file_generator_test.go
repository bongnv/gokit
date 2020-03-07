package command

import (
	"bytes"
	"testing"

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
	generator := &fileGenerator{
		filePath:     "main.go",
		templateName: "main",
		service: &Service{
			Package:     "github.com/hello",
			PackageName: "hello",
		},
		writer: mockWriter,
	}

	err := generator.Do()
	require.NoError(t, err)
	require.Equal(t,
		`main.go:
package main

import (
	"log"

	gokitServer "github.com/bongnv/gokit/util/server"
	"github.com/hello/internal/handlers"
	"github.com/hello/internal/server"
)

func main() {
	opts := []gokitServer.Option{}
	opts = append(opts, server.GetServiceOptions(handlers.New())...)

	err := gokitServer.Serve(
		opts...,
	)

	log.Println("Service stopped with:", err)
}
`,
		mockWriter.b.String(),
	)
}
