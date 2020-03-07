package command

import (
	"testing"

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
