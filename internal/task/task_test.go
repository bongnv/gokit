package task

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type taskFunc func() error

func (t taskFunc) Do() error {
	return t()
}

func Test_taskGroup(t *testing.T) {
	val := 0
	mockTask := func() error {
		val++
		return nil
	}

	mockTaskErr := func() error {
		return errors.New("random error")
	}

	tasks := Group{
		taskFunc(mockTask),
		taskFunc(mockTask),
		taskFunc(mockTaskErr),
		taskFunc(mockTask),
	}

	err := tasks.Do()
	require.Error(t, err)
	require.Equal(t, 2, val)
}
