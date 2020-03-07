package command

import (
	"testing"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/writer"
	"github.com/google/subcommands"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_getFilePath(t *testing.T) {
	c := &genCmd{
		path:          "root",
		interfaceName: "Service",
	}

	filePath := c.getFilePath("endpoint", "endpoints")
	require.Equal(t, "root/internal/endpoint/z_service_endpoints.go", filePath)
}

func Test_genCmd_Execute(t *testing.T) {
	mockParser := &parser.MockParser{}
	mockWriter := &writer.MockWriter{}
	cmd := &genCmd{
		parser:        mockParser,
		writer:        mockWriter,
		path:          ".",
		interfaceName: "Service",
	}

	mockParser.On("Parse").Return(&parser.Service{
		Name:        "Service",
		Package:     "github.com/hello",
		PackageName: "hello",
	}, nil).Once()
	mockWriter.On("Write", "internal/endpoint/z_service_endpoints.go", mock.Anything).Return(nil).Once()
	mockWriter.On("Write", "internal/server/z_service_server.go", mock.Anything).Return(nil).Once()

	resp := cmd.Execute(nil, nil)
	require.Equal(t, subcommands.ExitSuccess, resp)
	mockParser.AssertExpectations(t)
	mockWriter.AssertExpectations(t)
}
