package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"server/internal/repository"
	"server/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
	DB     repository.DatabaseRepo
}

func main() {
	log.Println("Starting server on port", port)

	var app application

	flag.StringVar(&app.DSN, "dns", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres for the application")
	flag.Parse()

	connect, err := app.connectDb()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: connect}
	defer app.DB.ConnectDb().Close()

	app.Domain = "http://localhost:8080"

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
