package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get any parameters
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update existing book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get any parameters
	// Loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...)	
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = params["id"]
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get any parameters
	// Loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...)	
		break
		}
	}
	json.NewEncoder(w).Encode(books)
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
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}