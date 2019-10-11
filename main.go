package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Book struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	PublishDate string `json:"publish_date"`
}

func handleBaseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my bookstore")
	fmt.Println("Hit the base route")
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	log.Println("route posted: /createNewBook")
	fmt.Println(r.Body)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading the body in createNewBook: %s", err.Error())
	}
	var newBook Book
	json.Unmarshal(reqBody, &newBook)
	db.Create(&newBook)

	json.NewEncoder(w).Encode(newBook)
}

func returnBooks(w http.ResponseWriter, r *http.Request) {
	books := []Book{}
	db.Find(&books)

	json.NewEncoder(w).Encode(books)
}

func returnOneBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var book Book
	db.Find(&book, id)

	fmt.Println(book)
	json.NewEncoder(w).Encode(book)
}

func updateOneBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updateBook Book
	json.Unmarshal(reqBody, &updateBook)
	db.Save(&updateBook)

	json.NewEncoder(w).Encode(updateBook)
}

func handleRequests() {
	log.Println("Starting development server 127.0.0.1:9000")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", handleBaseRoute)
	myRouter.HandleFunc("/create-book", createNewBook).Methods("POST")
	myRouter.HandleFunc("/books/", returnBooks).Methods("GET")
	myRouter.HandleFunc("/books/{id}", returnOneBook).Methods("GET")
	myRouter.HandleFunc("/update-book", updateOneBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func main() {
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=redeamapi dbname=book_store password=N+JmM7za4^zvq4ezK-dcc*dbszRWQ*9fDc$W9Ud sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening postgres instance: %s", err)
	}
	log.Println("Connection to database established successfully.")
	defer db.Close()

	db.AutoMigrate(&Book{})
	handleRequests()
}
