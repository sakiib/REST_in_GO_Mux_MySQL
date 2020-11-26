package db

import (
	"database/sql"
	"fmt"
)

// CreateDatabase ...
func CreateDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "sakibalamin:1620801042@/book_db")
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
