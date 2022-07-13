package db

import "database/sql"

type DBHelper interface {
	Open() *sql.DB
	Close() error
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}
