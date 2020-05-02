package db

import (
	"database/sql"

	// postgres driver
	_ "github.com/lib/pq"
)

// DB instance
type DB struct {
	Client *sql.DB
}

// Get establishes connection to database and returns pointer to DB instance
func Get(connStr string) (*DB, error) {
	db, err := get(connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

// Close func flose the DB connection
func (db *DB) Close() error {
	return db.Client.Close()
}

func get(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
