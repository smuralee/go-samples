package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

type FixedResponse struct {
	Status        string `json:"Status"`
	RemoteAddress string `json:"RemoteAddr"`
	Hostname      string `json:"Hostname"`
}

var Books []Book

var response FixedResponse

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response.Status = "Success"
	response.RemoteAddress = r.RemoteAddr
	response.Hostname, _ = os.Hostname()

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Printf("Fetching all the books...\n")

	err := json.NewEncoder(w).Encode(Books)
	if err != nil {
		panic(err)
	}
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("Fetching a book by id - %s ...\n", id)

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

	fmt.Printf("Adding a book...\n")

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

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("Updating a book using the id - %s ...\n", id)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedBook Book

	marshalErr := json.Unmarshal(reqBody, &updatedBook)
	if marshalErr != nil {
		panic(marshalErr)
	}

	for index, book := range Books {
		if book.Id == id {
			fmt.Printf("Match found for a book update, updating...\n")
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

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("Deleting a book by id - %s ...\n", id)

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
		{Id: "3", Title: "The Feynman Lectures on Physics", Author: "Richard P. Feynman"},
	}
	handleRequests()
}
