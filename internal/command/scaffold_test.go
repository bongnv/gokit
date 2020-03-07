package command

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_scaffoldCmd(t *testing.T) {
	t.Run("empty-package", func(t *testing.T) {
		c := &scaffoldCmd{}
		err := c.Do()
		require.Error(t, err)
	})
}
