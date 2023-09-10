package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jamieheart/go-books-mysql/pkg/utils"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = nil
var _ mysql.MySQLDriver

func Connect() {
	d, err := gorm.Open("mysql", utils.SQL_CONNECTION)
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
