package app

import (
	"github.com/jinzhu/gorm"
	"gogs.ballantine.tech/gballan1/gowis/models"
	// loading SQLite dialect for gorm ORM
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDB - sets up the database
func InitDB() *gorm.DB {
	// setup our DB
	db, dbErr := gorm.Open("sqlite3", "gowis.db")

	// check for DB errors
	if dbErr != nil {
		panic(dbErr)
	}

	// do some migrations
	db.AutoMigrate(&models.Page{})

	// return the DB object
	return db
}
