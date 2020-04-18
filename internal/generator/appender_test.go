package generator

import (
	"testing"

	"github.com/bongnv/gokit/internal/iohelper"
	"github.com/bongnv/gokit/internal/parser"
	"github.com/stretchr/testify/require"
)

func TestAppender_Do(t *testing.T) {
	mockWriter := &bufWriter{}
	mockReader := &iohelper.MockReader{}
	generator := &Appender{
		FilePath:     "main.go",
		TemplateName: "entity",
		Data: &parser.Data{
			Name: "Entity",
		},
		Writer: mockWriter,
		Reader: mockReader,
	}

	mockReader.On("Read", "main.go").Return([]byte(`package main`), nil).Once()
	err := generator.Do()
	require.NoError(t, err)
	require.Equal(t,
		`main.go:
package main

import "time"

// Entity ...
//go:generate gokitgen entity -name Entity
type Entity struct {
	ID        int64     `+"`"+`gorm:"primary_key"`+"`"+`
	CreatedAt time.Time `+"`"+`gorm:"not null"`+"`"+`
	UpdatedAt time.Time `+"`"+`gorm:"not null"`+"`"+`
}
`,
		mockWriter.b.String(),
	)
	mockReader.AssertExpectations(t)
}
