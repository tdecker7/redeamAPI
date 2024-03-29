package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Book struct {
	Id          string `sql:"type:serial PRIMARY KEY"`
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
	log.Println("route posted: /books")
	books := []Book{}
	db.Find(&books)

	json.NewEncoder(w).Encode(books)
}

func returnOneBook(w http.ResponseWriter, r *http.Request) {
	log.Println("route posted: /books/{id}")
	vars := mux.Vars(r)
	id := vars["id"]
	var book Book
	db.Find(&book, id)

	fmt.Println(book)
	json.NewEncoder(w).Encode(book)
}

func updateOneBook(w http.ResponseWriter, r *http.Request) {
	log.Println("route posted: /update-book/{id}")
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updateBook Book
	updateBook.Id = id
	json.Unmarshal(reqBody, &updateBook)
	db.Save(&updateBook)

	json.NewEncoder(w).Encode(updateBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("route posted: /delete-book/{id}")
	vars := mux.Vars(r)
	id := vars["id"]

	bookToDelete := Book{Id: id}
	db.Find(&bookToDelete, id)
	db.Delete(&bookToDelete)
	json.NewEncoder(w).Encode(bookToDelete)
}

func handleRequests() {
	log.Println("Starting development server 127.0.0.1:9000")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", handleBaseRoute)
	myRouter.HandleFunc("/create-book", createNewBook).Methods("POST")
	myRouter.HandleFunc("/books/", returnBooks).Methods("GET")
	myRouter.HandleFunc("/books/{id}", returnOneBook).Methods("GET")
	myRouter.HandleFunc("/update-book/{id}", updateOneBook).Methods("POST")
	myRouter.HandleFunc("/delete-book/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func main() {
	for i := 0; i < 10; i++ {
		db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=book_store sslmode=disable")
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	
	if err != nil {
		log.Fatalf("Error opening postgres insftance: %s", err)
	}
	log.Println("Connection to database established successfully.")
	defer db.Close()

	db.AutoMigrate(&Book{})
	handleRequests()
}
