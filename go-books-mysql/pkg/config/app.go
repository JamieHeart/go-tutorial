package config

import (
	"github.com/jamieheart/go-books-mysql/pkg/utils"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = nil

func Connect() {
	d, err := gorm.Open(utils.SQL_USERNAME, utils.SQL_CONNECTION)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	panic("db is nil!")
}
