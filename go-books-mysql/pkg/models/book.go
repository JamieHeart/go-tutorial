package models

import (
	"errors"

	"github.com/jamieheart/go-books-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model         //	why is the id retruning wrong via the api, but correct in the db?
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Isbn        string `json:"isbn" gorm:"unique"`
}

func Init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	db.NewRecord(b)
	rows := db.Create(&b)
	if rows.RowsAffected == 0 {
		return b, errors.New("duplicate entry")
	}
	return b, nil
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id string) (Book, error) {
	var book Book
	rows := db.Find(&book, id)
	if rows.RowsAffected == 0 {
		return book, errors.New("not found")
	}
	return book, nil
}
