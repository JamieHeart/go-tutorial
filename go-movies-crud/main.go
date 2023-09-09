package main

import (
	//	Will be printing out
	//	Will be reporting errors
	//	Json input and responses
	//	Generate some entropy
	//	It is a server, after all

	"github.com/gorilla/mux" //	advanced GO router
)

type Movie struct {
	ID       string    `json:"id"` //	`json:"[how to serialize/deserialize]"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies() {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies)
}
