package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// CreateDatabase ...
func CreateDatabase() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("dbUsername")
	dbPassword := os.Getenv("dbPassword")

	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/book_db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("DB connected!")
	return db, nil
}
