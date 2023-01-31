package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func slowQuery(ctx context.Context) error {
	// Create a new child context with a 5-second timeout, using the
	// provided ctx parameter as the parent.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Pass the child context (the one with the timeout) as the first
	// parameter to ExecContext().
	_, err := db.ExecContext(ctx, "SELECT pg_sleep(10)")
	return err
}

func main() {
	var err error

	db, err = sql.Open("postgres", "postgres://postgres:123@localhost/example_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", exampleHandler)

	log.Print("Listening...")
	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	err := slowQuery(r.Context())
	if err != nil {
		serverError(w, err)
		return
	}

	fmt.Fprintln(w, "OK")
}

func serverError(w http.ResponseWriter, err error) {
	log.Printf("ERROR: %s", err.Error())
	http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
}
