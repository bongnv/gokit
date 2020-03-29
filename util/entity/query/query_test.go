package query

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

func Test_Transform(t *testing.T) {
	expectedDB := &gorm.DB{}
	mockQuery := Query(func(db *gorm.DB) *gorm.DB {
		return expectedDB
	})

	newDB := Transform(nil, mockQuery)
	require.Equal(t, expectedDB, newDB)
}
