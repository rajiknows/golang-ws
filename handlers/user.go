package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rajiknows/vedashala/config"
	_ "github.com/rajiknows/vedashala/config"
	"github.com/rajiknows/vedashala/internal/database"
	"github.com/rajiknows/vedashala/utils"
	"net/http"
	"time"
)

type tokens struct {
	accessToken   string
	refresh_token string
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	subrouter := chi.NewRouter()
	subrouter.Use()
	subrouter.Post("/login", handleLogin)
	subrouter.Post("/register", handleRegister)
	// ... other user routes

	subrouter.ServeHTTP(w, r)

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		AccessToken string `json:"access_token"`
	}

	var body Body
	cookie := r.Cookies()
	println(cookie)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	// For demonstration purposes, we're just echoing the access token back
	response := fmt.Sprintf("Received access token: %s", body.AccessToken)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	db := config.GetConfig().DB
	var body RegisterBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.RespondWithError(w, 404, "unsupported data ")
		return
	}

	password, err := utils.HashPassword(body.Password)
	if err != nil {
		fmt.Print("password hashing failed")
	}

	user, err := db.GetUserByEmail(r.Context(), body.Email)
	if err != nil {
		user, err := db.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      body.Email,
			Email:     body.Email,
			Password:  password,
		})

		if err != nil {
			utils.RespondWithError(w, 500, "couldnt create a user")
		}
		w.WriteHeader(200)
		utils.RespondWithJson(w, 200, utils.DatabaseUserToUser(user))

	}

	utils.RespondWithError(w, 400, fmt.Sprint("%s email already exists", user.Email))
	http.Redirect(w, r, "/login", 400)

	//_ := config.GetConfig().DB
	// Registration logic here using db

}

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
