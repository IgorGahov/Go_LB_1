package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	connStr := "user=postgres password=sfhdkfjh233KKKdssas89 dbname=webapp sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
