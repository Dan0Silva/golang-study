package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // sql driver connection
)

// open database connection
func Connect() (*sql.DB, error) {
	stringConnection := "danilo:2204@tcp(172.17.0.2:3306)/golang_study?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
