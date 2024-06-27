package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB() *gorm.DB {
	if DB == nil {
		database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		DB = database
	}

	return DB
}
