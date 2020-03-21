package generator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_helper_funcs(t *testing.T) {
	require.Equal(t, "todos", toPlural("todo"))
	require.Equal(t, "todo", toLower("Todo"))
	require.Equal(t, "&Todo{}", initValue("*Todo"))
}
