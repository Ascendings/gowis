package models

import (
	"github.com/jinzhu/gorm"
	// loading SQLite dialect for gorm ORM
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// DB - gorm database engine object
	DB *gorm.DB
)

// InitDB - sets up the database
func InitDB() {
	// setup our DB
	DB, dbErr := gorm.Open("sqlite3", "gowis.db")

	// check for DB errors
	if dbErr != nil {
		panic(dbErr)
	}

	// do some migrations
	DB.AutoMigrate(&Page{})
}
