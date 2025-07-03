package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"server/internal/repository"
	"server/internal/repository/dbrepo"
	"time"
)

const port = 8080

type application struct {
	Domain       string
	DSN          string
	DB           repository.DatabaseRepo
	Auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	log.Println("Starting server on port", port)

	var app application

	flag.StringVar(&app.DSN, "dns", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres for the application")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "test jwt secret", "signing secret for JWT tokens")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "test jwt issuer", "signing issuer for JWT tokens")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "test jwt audience", "signing audience for JWT tokens")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "localhost", "domain")
	flag.Parse()

	connect, err := app.connectDb()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: connect}
	defer app.DB.ConnectDb().Close()

	app.Auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   15 * time.Minute,
		RefreshExpiry: 24 * time.Hour,
		CookieDomain:  app.CookieDomain,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
