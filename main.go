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
	fmt.Println("server running at port :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, app.Router))
}
