package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jimyag/bookstore/models"
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

	env := &Env{db: db}

	// Pass the Env struct as a parameter to booksIndex().
	http.HandleFunc("/books", booksIndex(env))
	http.ListenAndServe(":3000", nil)
}

// Use a closure to make Env available to the handler logic.
func booksIndex(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
