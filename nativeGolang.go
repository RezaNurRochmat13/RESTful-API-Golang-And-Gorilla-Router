package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/booksCategory", BooksCategory).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome in Index")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Books Index")
}

func BooksCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Books Category")
}
