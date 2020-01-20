package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:firstname`
	Lastname  string `json:firstname`
}

// create slice of book = books array
var books []Book

// Get all books

func getBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(books)
}

// Get signle bool

func getBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	params := mux.Vars(req)
	// loop and find id
	for _, book := range books {
		if params["id"] == book.ID {
			json.NewEncoder(res).Encode(book)
			return
		}
	}
	// return couldnot find book id
	json.NewEncoder(res).Encode("Could not find book id")
}

// Create New Book

func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000)) // mockId = NOT SAFE
	books = append(books, book)
	json.NewEncoder(res).Encode(book)
}

// update a Book

func updateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var updatedBook Book
	_ = json.NewDecoder(req.Body).Decode(&updatedBook)
	params := mux.Vars(req)
	for index, book := range books {
		if params["id"] == book.ID {
			updatedBook.ID = book.ID
			books[index] = updatedBook
			break
		}
	}
	json.NewEncoder(res).Encode(updatedBook)
}

// Delete a Book

func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	params := mux.Vars(req)
	for index, book := range books {
		if params["id"] == book.ID {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "214312f", Title: "Desgin Patterns", Author: &Author{Firstname: "Gang", Lastname: "Of4"}})
	books = append(books, Book{ID: "2", Isbn: "321", Title: "Desgin IDEAs", Author: &Author{Firstname: "Gang", Lastname: "Of 3"}})
	// Route Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
