package models

import "database/sql"

// This time the global variable is unexported.
var db *sql.DB

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// InitDB sets up the connection pool global variable.
func InitDB(dataSourceName string) error {
	var err error

	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	return db.Ping()
}

// AllBooks returns a slice of all books in the books table.
func AllBooks() ([]Book, error) {
	// Note that we are calling Query() on the global variable.
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
