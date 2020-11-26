package main

import (
	"fmt"
	"log"
	"net/http"
	"restapisql/app"
	"restapisql/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	app.HandleRequests()
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
