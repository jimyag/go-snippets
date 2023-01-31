package models

import (
	"context"
	"database/sql"
	"errors"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks(ctx context.Context) ([]Book, error) {
	// Retrieve the connection pool from the context. Because the
	// r.Context().Value() method always returns an interface{} type, we
	// need to type assert it into a *sql.DB before using it.
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return nil, errors.New("could not get database connection pool from context")
	}
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []Book

	for rows.Next() {
		var bk Book

		err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
