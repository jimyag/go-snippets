package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/darjun/bookstore/models"
)

// Create a custom Env struct which holds a connection pool.
type Env struct {
	db *sql.DB
}

func main() {
	// Initialise the connection pool.
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of Env containing the connection pool.
	env := &Env{db: db}

	// Use env.booksIndex as the handler function for the /books route.
	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

// Define booksIndex as a method on Env.
func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	bks, err := models.AllBooks(env.db)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
