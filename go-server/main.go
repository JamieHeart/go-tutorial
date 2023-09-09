package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT = "8888"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found: "+r.URL.Path, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method "+r.Method+" not supported", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found: "+r.URL.Path, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method "+r.Method+" not supported", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "POST success")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Print("Starting server on port " + PORT)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
