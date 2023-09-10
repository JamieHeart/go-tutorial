package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamieheart/go-books-mysql/pkg/models"
	"github.com/jamieheart/go-books-mysql/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, &newBook)
	b, err := newBook.CreateBook()
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	res, _ := json.Marshal(b)

	w.Header().Add("Content-Type", "application/json")
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

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := models.GetBookById(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
