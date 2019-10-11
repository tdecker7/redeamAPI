package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publish_date"`
}

func handleBaseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my bookstore")
	fmt.Println("Hit the base route")
}

func handleRequests() {
	log.Println("Starting development server 127.0.0.1:9000")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", handleBaseRoute)
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=redeamapi dbname=book_store password=N+JmM7za4^zvq4ezK-dcc*dbszRWQ*9fDc$W9Ud sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening postgres instance: %s", err)
	}
	log.Println("Connection to database established successfully.")
	defer db.Close()

	db.AutoMigrate(&Book{})
	handleRequests()
}
