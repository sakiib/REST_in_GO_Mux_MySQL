package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restapisql/app"
	"restapisql/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World!")

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	app := &app.App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	app.HandleRequests()
	log.Fatal(http.ListenAndServe(":"+PORT, app.Router))
}
