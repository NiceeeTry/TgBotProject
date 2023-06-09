package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(driver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
