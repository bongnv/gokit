package command

import (
	"testing"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/writer"
	"github.com/google/subcommands"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_entityCmd_Execute(t *testing.T) {
	mockParser := &parser.MockParser{}
	mockWriter := &writer.MockWriter{}
	cmd := &entityCmd{
		entityParser: mockParser,
		writer:       mockWriter,
		path:         ".",
		entityName:   "Entity",
	}

	mockParser.On("Parse", ".", "Entity").Return(&parser.Data{
		Name:        "Service",
		Package:     "github.com/hello",
		PackageName: "hello",
	}, nil).Once()
	mockWriter.On("Write", "z_entity_dao.go", mock.Anything).Return(nil).Once()

	resp := cmd.Execute(nil, nil)
	require.Equal(t, subcommands.ExitSuccess, resp)
	mockParser.AssertExpectations(t)
	mockWriter.AssertExpectations(t)
}
