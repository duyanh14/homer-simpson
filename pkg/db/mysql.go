package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dbHelper struct {
	db *sql.DB
}

// dsn connect to db
func dsn(dbName, username, password, hostname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func InitMysqlDB(host string, port int, username, passsword, database string) (*sql.DB, error) {
	hostPort := fmt.Sprintf("%v:%d", host, port)
	connectString := dsn(database, username, passsword, hostPort)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}

func NewMysqlDBHelper(host string, port int, username, password, database string) DBHelper {
	db, err := InitMysqlDB(host, port, username, password, database)

	if err != nil {
		panic(err)
	}
	return &dbHelper{
		db: db,
	}
}

func (h *dbHelper) Open() *sql.DB {
	return h.db
}

func (h *dbHelper) Close() error {
	return h.db.Close()
}

func (h *dbHelper) Begin() (*sql.Tx, error) {
	return h.db.Begin()
}

func (h *dbHelper) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (h *dbHelper) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
