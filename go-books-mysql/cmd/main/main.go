package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamieheart/go-books-mysql/pkg/routes"
	"github.com/jamieheart/go-books-mysql/pkg/utils"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(utils.HTTP_SERVER_URI, r))
}
