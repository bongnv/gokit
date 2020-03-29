package query

import (
	"github.com/jinzhu/gorm"
)

// Query is a flexible pattern to allow query DB.
type Query func(db *gorm.DB) *gorm.DB

// Transform applies multiple query to an existing instance of gorm.DB to create a new gorm.DB.
func Transform(db *gorm.DB, queries ...Query) *gorm.DB {
	for _, q := range queries {
		db = q(db)
	}

	return db
}
