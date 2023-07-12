package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

var db *sql.DB

// InitializeDB initializes the database connection
func InitializeDB() error {
	// Open the database connection

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Successfully connected!")

	conn, err := sql.Open("mysql", "user:password@tcp(hostname:port)/database")
	if err != nil {
		return err
	}

	// Assign the database connection to the global variable
	db = conn
	return nil
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
