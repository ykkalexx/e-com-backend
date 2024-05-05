/*
	e-com-backend is a simple e-commerce backend application that provides a RESTful API for managing products and orders.
	The application is built using the Go programming language and the MySQL database.
*/

package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/ykkalexx/e-com-backend/cmd/api"
	"github.com/ykkalexx/e-com-backend/config"
	"github.com/ykkalexx/e-com-backend/db"
)

/*
	The main function is the entry point of the application.
	The function sets up the database connection and the API server.
*/

func main() {
	// setting up the database connection
	db, err := db.NewMySQL(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPass,
		Addr: config.Envs.DBAddr,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	// setting up the API server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// start database connection
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful")
}