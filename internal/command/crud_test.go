package command

import (
	"testing"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/iohelper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_crudCmd(t *testing.T) {
	mockParser := &parser.MockParser{}
	mockCrudParser := &parser.MockParser{}
	mockEntityParser := &parser.MockParser{}
	mockWriter := &iohelper.MockWriter{}
	mockReader := &iohelper.MockReader{}
	c := &crudCmd{
		path:          ".",
		resource:      "Todo",
		serviceParser: mockParser,
		crudParser:    mockCrudParser,
		entityParser:  mockEntityParser,
		writer:        mockWriter,
		reader:        mockReader,
	}

	mockReader.On("Read", "internal/storage/storage.go").Return([]byte("package storage"), nil).Once()
	mockWriter.On("Write", "todo_service.go", mock.Anything).Once().Return(nil)
	mockCrudParser.On("Parse", ".", "Todo").Once().Return(&parser.Data{
		Name:        "Todo",
		Package:     "github.com/todo",
		PackageName: "todo",
	}, nil)
	mockParser.On("Parse", ".", "TodoService").Once().Return(&parser.Data{
		Name:        "TodoService",
		Package:     "github.com/todo",
		PackageName: "todo",
	}, nil)
	mockWriter.On("Write", "internal/todoservice/z_endpoints.go", mock.Anything).Once().Return(nil)
	mockWriter.On("Write", "internal/todoservice/z_server.go", mock.Anything).Once().Return(nil)
	mockWriter.On("Write", "internal/handlers/todo_handler.go", mock.Anything).Once().Return(nil)
	mockWriter.On("Write", "internal/storage/storage.go", mock.Anything).Once().Return(nil)
	mockWriter.On("Write", "internal/storage/z_todo_dao.go", mock.Anything).Return(nil).Once()
	mockEntityParser.On("Parse", "internal/storage", "Todo").Return(&parser.Data{
		Name:        "Todo",
		Package:     "github.com/todo/internal/storage",
		PackageName: "storage",
	}, nil).Once()

	err := c.Do()
	require.NoError(t, err)
	mockParser.AssertExpectations(t)
	mockWriter.AssertExpectations(t)
	mockReader.AssertExpectations(t)
	mockEntityParser.AssertExpectations(t)
}
