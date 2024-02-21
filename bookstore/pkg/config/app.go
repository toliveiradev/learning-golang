package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB() {
	database, err := gorm.Open("mysql", "bookstore.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
