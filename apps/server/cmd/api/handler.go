package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
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

func (app *application) RefreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.Auth.CookieName {
			claims := &Claims{}
			refToken := cookie.Value
			if _, err := jwt.ParseWithClaims(refToken, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(app.JWTSecret), nil
			}); err != nil {
				app.errorJson(w, errors.New("invalid refresh token"), http.StatusUnauthorized)
				return
			}

			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJson(w, errors.New("user not found"), http.StatusNotFound)
				return
			}
			user, err := app.DB.GetUserById(userID)
			if err != nil {
				app.errorJson(w, errors.New("user not found"), http.StatusNotFound)
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
			app.writeJson(w, http.StatusOK, tokenPair)
		}
	}
}
