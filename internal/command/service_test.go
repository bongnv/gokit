package command

import (
	"testing"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/iohelper"
	"github.com/google/subcommands"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_getFilePath(t *testing.T) {
	c := &serviceCmd{
		path:          "root",
		interfaceName: "Service",
	}

	filePath := c.getFilePath("endpoints")
	require.Equal(t, "root/internal/service/z_endpoints.go", filePath)
}

func Test_serviceCmd_Execute(t *testing.T) {
	mockParser := &parser.MockParser{}
	mockWriter := &iohelper.MockWriter{}
	cmd := &serviceCmd{
		parser:        mockParser,
		writer:        mockWriter,
		path:          ".",
		interfaceName: "Service",
	}

	mockParser.On("Parse", ".", "Service").Return(&parser.Data{
		Name:        "Service",
		Package:     "github.com/hello",
		PackageName: "hello",
	}, nil).Once()
	mockWriter.On("Write", "internal/service/z_endpoints.go", mock.Anything).Return(nil).Once()
	mockWriter.On("Write", "internal/service/z_server.go", mock.Anything).Return(nil).Once()

	resp := cmd.Execute(nil, nil)
	require.Equal(t, subcommands.ExitSuccess, resp)
	mockParser.AssertExpectations(t)
	mockWriter.AssertExpectations(t)
}
