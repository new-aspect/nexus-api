package model

import (
	"github.com/new-aspect/nexus-api/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	DB, err = gorm.Open(sqlite.Open("nexus-api.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(&Channel{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Token{}); err != nil {
		return err
	}

	common.UsingSQLite = true

	return
}
