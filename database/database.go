package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}
