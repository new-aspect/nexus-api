package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = gorm.Open(sqlite.Open("nexus-api.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&Channel{})
}
