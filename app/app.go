package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App ...
type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := app.Database.Query("select * from Book")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var books []Book
	b := &Book{}
	for rows.Next() {
		err := rows.Scan(&b.Name, &b.ID)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, Book{Name: b.Name, ID: b.ID})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books)
}

// HandleRequests ...
func (app *App) HandleRequests() {
	app.Router.HandleFunc("/api/books", app.getBooks).Methods("GET")
}
