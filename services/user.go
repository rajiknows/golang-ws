package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/rajiknows/vedashala/config"
)

type tokens struct {
	accessToken   string
	refresh_token string
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	subrouter := chi.NewRouter()
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

	// Read the request body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse the request bodyBytes
	err = json.Unmarshal(bodyBytes, &body)
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
	//_ := config.GetConfig().DB
	// Registration logic here using db
	w.WriteHeader(200)
	w.Write([]byte("register"))
}
