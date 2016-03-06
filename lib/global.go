package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data/go-ci.db")
	if err != nil {
		log.Fatal("init database error, ", err)
		return nil
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return &db
}
