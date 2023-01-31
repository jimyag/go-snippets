package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darjun/bookstore/models"
)

func main() {
	// Use the InitDB function to initialise the global variable.
	err := models.InitDB("postgres://postgres:123@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}

// booksIndex send an HTTP response listing all books.
func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
