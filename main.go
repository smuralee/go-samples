package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

var Books []Book

func welcome(w http.ResponseWriter, r *http.Request) {
	bytes, err := fmt.Fprintf(w, "Welcome to the Library!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d Bytes written.\n", bytes)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Fetch all books")

	err := json.NewEncoder(w).Encode(Books)
	if err != nil {
		panic(err)
	}
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Fetch book by Id")

	vars := mux.Vars(r)
	id := vars["id"]

	for _, book := range Books {
		if book.Id == id {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				panic(err)
			}
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Add a book")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book

	marshalErr := json.Unmarshal(reqBody, &book)
	if marshalErr != nil {
		panic(marshalErr)
	}
	Books = append(Books, book)

	encodeErr := json.NewEncoder(w).Encode(book)
	if encodeErr != nil {
		panic(encodeErr)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Update a book")

	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedBook Book

	marshalErr := json.Unmarshal(reqBody, &updatedBook)
	if marshalErr != nil {
		panic(marshalErr)
	}

	for index, book := range Books {
		if book.Id == id {
			fmt.Println("Match found for a book update")
			Books = append(Books[:index], Books[index+1:]...)
			Books = append(Books, updatedBook)
			encodeErr := json.NewEncoder(w).Encode(updatedBook)
			if encodeErr != nil {
				panic(encodeErr)
			}
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Delete a book")

	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			fmt.Println("Match found for a book delete")
			Books = append(Books[:index], Books[index+1:]...)
			return
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", welcome)
	myRouter.HandleFunc("/books", getAllBooks).Methods("GET")
	myRouter.HandleFunc("/books/{id}", getBookById).Methods("GET")
	myRouter.HandleFunc("/books", createBook).Methods("POST")
	myRouter.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	myRouter.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Books = []Book{
		{Id: "1", Title: "Mozart in the Jungle", Author: "Blair Tindall"},
		{Id: "2", Title: "Bad Blood", Author: "John Carreyrou"},
	}
	handleRequests()
}
