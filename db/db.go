package db

/*
	The NewMySQL function is a constructor function that returns a new MySQL database connection.
	The function takes a mysql.Config struct as an argument and returns a pointer to a sql.DB struct.
*/

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQL(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}