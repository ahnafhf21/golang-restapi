package main

import(
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)
// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books variable as a slice Book Struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request){

}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request){

}

// Update existing book
func updateBook(w http.ResponseWriter, r *http.Request){

}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){

}

func main(){
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "1945", Title: "Hujan Bulan Juni", Author: &Author{
		Firstname: "Sapardi", Lastname: "Joko Darmono", }})
	books = append(books, Book{ID: "2", Isbn: "8291389", Title: "Rembulan Tenggelam Di Wajahmu", Author: &Author{
		Firstname: "Tere", Lastname: "Liye", }})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}