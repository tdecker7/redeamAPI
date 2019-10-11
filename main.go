package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=redeamapi dbname=book_store password=N+JmM7za4^zvq4ezK-dcc*dbszRWQ*9fDc$W9Ud sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening postgres instance: %s", err)
	}
	defer db.Close()
}
