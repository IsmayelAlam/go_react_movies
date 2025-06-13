package main

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/models"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var response = struct {
		Message string `json:"message"`
	}{
		Message: "Welcome to the API",
	}

	output, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	var movies []models.Movie

	inception := models.Movie{
		Title:       "Inception",
		Description: "A mind-bending thriller by Christopher Nolan.",
		ID:          1,
		ReleaseDate: time.Now(),
		RunTime:     148,
		MpaaRating:  "PG-13",
		Image:       "http://example.com/inception.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	interstellar := models.Movie{
		Title:       "Interstellar",
		Description: "A journey through space and time to save humanity.",
		ID:          2,
		ReleaseDate: time.Now(),
		RunTime:     169,
		MpaaRating:  "PG-13",
		Image:       "http://example.com/interstellar.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	dunkirk := models.Movie{
		Title:       "Dunkirk",
		Description: "A gripping war epic showcasing survival against the odds.",
		ID:          3,
		ReleaseDate: time.Now(),
		RunTime:     106,
		MpaaRating:  "PG-13",
		Image:       "http://example.com/dunkirk.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, inception, interstellar, dunkirk)

	output, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
