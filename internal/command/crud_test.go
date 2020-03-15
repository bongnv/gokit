package command

import (
	"testing"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/writer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_crudCmd(t *testing.T) {
	mockParser := &parser.MockParser{}
	mockWriter := &writer.MockWriter{}
	c := &crudCmd{
		path:     ".",
		resource: "Todo",
		parser:   mockParser,
		writer:   mockWriter,
	}

	mockWriter.On("Write", "z_todo.go", mock.Anything).Once().Return(nil)
	mockParser.On("Parse", ".", "TodoService").Once().Return(&parser.Service{
		Name:        "TodoService",
		Package:     "github.com/todo",
		PackageName: "todo",
	}, nil)
	mockWriter.On("Write", "internal/todoservice/z_endpoints.go", mock.Anything).Once().Return(nil)
	mockWriter.On("Write", "internal/todoservice/z_server.go", mock.Anything).Once().Return(nil)

	err := c.Do()
	require.NoError(t, err)
	mockParser.AssertExpectations(t)
	mockWriter.AssertExpectations(t)
}
