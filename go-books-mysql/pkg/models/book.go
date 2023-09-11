package models

import (
	"errors"
	"strconv"

	"github.com/jamieheart/go-books-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
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

func (b *Book) CreateBook() (*Book, error) { //	TODO	:	https://github.com/JamieHeart/go-tutorial/issues/6 - Losing PK Space on conflict
	db.NewRecord(b)
	rows := db.Create(&b)
	if rows.RowsAffected == 0 { //	TODO: https://github.com/JamieHeart/go-tutorial/issues/5
		return b, errors.New("duplicate entry") //	OOPS! This means if you delete a book you can not add it again
	}
	return b, nil
}

func DeleteBook(id string) error {
	book, err := GetBookById(id)
	if err != nil {
		return err
	}
	db.Delete(book)
	return nil
}

func (b *Book) UpdateBook() (*Book, error) {
	book, err := GetBookById(strconv.FormatUint(uint64(b.ID), 10))
	if err != nil {
		return &book, err
	}

	book.Author = b.Author
	book.Isbn = b.Isbn
	book.Name = b.Name
	book.Publication = b.Publication

	db.Save(&book)

	return &book, nil
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
