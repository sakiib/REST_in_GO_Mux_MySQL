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

func (app *App) getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	book := &Book{}
	row := app.Database.QueryRow("select * from Book where id = ?", params["id"])
	err := row.Scan(&book.Name, &book.ID)

	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)
}

func (app *App) createBook(w http.ResponseWriter, r *http.Request) {

}

func (app *App) updateBook(w http.ResponseWriter, r *http.Request) {

}

func (app *App) deleteBook(w http.ResponseWriter, r *http.Request) {

}

// HandleRequests ...
func (app *App) HandleRequests() {
	app.Router.HandleFunc("/api/books", app.getBooks).Methods("GET")
	app.Router.HandleFunc("/api/books/{id}", app.getBook).Methods("GET")
	app.Router.HandleFunc("/api/books/", app.createBook).Methods("POST")
	app.Router.HandleFunc("/api/books/{id}", app.updateBook).Methods("PUT")
	app.Router.HandleFunc("/api/books/{id}", app.deleteBook).Methods("DELETE")
}
