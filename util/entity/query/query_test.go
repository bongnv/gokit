package query

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func Test_Equal(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&gorm.Model{})
	db.Create(&gorm.Model{ID: 1})

	var item gorm.Model
	result := Equal("id", 1)(db).Find(&item)
	require.NoError(t, result.Error)
	require.EqualValues(t, 1, item.ID)
}

func Test_Limit(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&gorm.Model{})
	db.Create(&gorm.Model{ID: 1})
	db.Create(&gorm.Model{ID: 2})

	var items []*gorm.Model
	result := Limit(1)(db).Find(&items)
	require.NoError(t, result.Error)
	require.Len(t, items, 1)
}

func Test_Offset(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&gorm.Model{})
	db.Create(&gorm.Model{ID: 1})
	db.Create(&gorm.Model{ID: 2})

	var items []*gorm.Model
	result := Limit(1)(Offset(1)(db)).Find(&items)
	require.NoError(t, result.Error)
	require.Len(t, items, 1)
	require.EqualValues(t, 2, items[0].ID)
}

func Test_OrderBy(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&gorm.Model{})
	db.Create(&gorm.Model{ID: 1})
	db.Create(&gorm.Model{ID: 2})

	var items []*gorm.Model
	result := OrderBy("id", false)(db).Find(&items)
	require.NoError(t, result.Error)
	require.Len(t, items, 1)
	require.EqualValues(t, 2, items[0].ID)
}
