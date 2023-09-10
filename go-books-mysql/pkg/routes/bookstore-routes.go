package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamieheart/go-books-mysql/pkg/controllers"
	"github.com/jamieheart/go-books-mysql/pkg/utils"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc(utils.ROUTE_BOOK, controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc(utils.ROUTE_BOOKS, controllers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc(utils.ROUTE_BOOK+"/{id}", controllers.GetBook).Methods(http.MethodGet)
}
