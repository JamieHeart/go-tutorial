package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jamieheart/go-books-mysql/pkg/models"
	"github.com/jamieheart/go-books-mysql/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, &newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)

	println(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
