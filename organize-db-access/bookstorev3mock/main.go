package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/darjun/bookstore/models"
)

type Env struct {
	// Replace the reference to models.BookModel with an interface
	// describing its methods instead. All the other code remains exactly
	// the same
	books interface {
		All() ([]models.Book, error)
	}
}

func main() {
	// Initialise the connection pool as normal.
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	// Initialise Env with a models.BookModel instance (which in turn wraps the connection pool).
	env := &Env{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

// Use a closure to make Env available to the handler logic.
func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	bks, err := env.books.All()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
