package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

var HTTP_PORT int = 8888

var MOVIES_ROUTE string = "/movies"
var MOVIE_ROUTE string = MOVIES_ROUTE + "/{id}"

var movies []Movie

func seedMovies() {
	movies = append(movies, Movie{ID: "0", Isbn: "1234567", Title: "Foo", Director: &Director{FirstName: "Bar", LastName: "Baz"}})
	movies = append(movies, Movie{ID: "1", Isbn: "7654321", Title: "Oof", Director: &Director{FirstName: "Rab", LastName: "Zab"}})
}
func setJsonResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	setJsonResponse(w)
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	setJsonResponse(w)
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	setJsonResponse(w)

	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(len(movies))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	setJsonResponse(w)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	setJsonResponse(w)
	params := mux.Vars(r)
	var match bool = false
	for index, item := range movies {
		if item.ID == params["id"] {
			match = true
			if index+1 == len(movies) {
				movies = movies[:len(movies)-1]
			} else {
				movies = append(movies[:index], movies[index+1])
			}
			break
		}
	}
	if match {
		json.NewEncoder(w).Encode(movies)
	} else {
		log.Println("not found")
		http.Error(w, "Movie not found", http.StatusNotFound)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc(MOVIES_ROUTE, getMovies).Methods(http.MethodGet)
	r.HandleFunc(MOVIE_ROUTE, getMovie).Methods(http.MethodGet)
	r.HandleFunc(MOVIES_ROUTE, createMovie).Methods(http.MethodPost)
	r.HandleFunc(MOVIE_ROUTE, updateMovie).Methods(http.MethodPut)
	r.HandleFunc(MOVIE_ROUTE, deleteMovie).Methods(http.MethodDelete)

	seedMovies()

	fmt.Printf("Server started on port %d", HTTP_PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(HTTP_PORT), r))
}
