# Simple Golang CRUD with docker-compose
This is a _very_ simple CRUD api built in golang. 

## Startup
To start the application run:

`$ docker-compose up`

The application starts on `http://localhost:9000`

## Routes
* Create: `/create-book/`, which takes a json body.
* Read: `/books/` will return all books. `/books/{id}/` to specify a book to return.
* Update: `/update-book/{id}`, which takes a json body.
* Delete: `/delete-book/{id}`

## Tests
Start the server separately in `docker-compose`, then run the tests which run against the server.
Run the following commands:
1. `$ docker-compose up -d`
1. `$ go test`


