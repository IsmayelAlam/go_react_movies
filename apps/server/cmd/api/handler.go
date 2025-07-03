package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
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
	movies, err := app.DB.AddMovies()

	if err != nil {
		app.errorJson(w, err)
		return
	}

	if err = app.writeJson(w, http.StatusOK, movies); err != nil {
		log.Fatal(err)
	}

}

func (app *application) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := app.readJson(w, r, &requestPayload); err != nil {
		app.errorJson(w, err)
		return
	}

	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	if valid, err := user.PasswordMatch(requestPayload.Password); err != nil || !valid {
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokenPair, err := app.Auth.GenerateToken(&u)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	refreshCookie := app.Auth.GetRefreshCookie(tokenPair.RefreshToken)

	http.SetCookie(w, refreshCookie)

	if err = app.writeJson(w, http.StatusAccepted, tokenPair); err != nil {
		log.Fatal(err)
	}
}
