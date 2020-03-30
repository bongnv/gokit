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

// Equal implements equal query.
func Equal(field string, value interface{}) Query {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+" = ?", value)
	}
}

// Limit implements limit query.
func Limit(limit int64) Query {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

// Offset implements offset query.
func Offset(offset int64) Query {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}
