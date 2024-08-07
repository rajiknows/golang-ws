package services

import (
	"net/http"

	_ "github.com/rajiknows/vedashala/config"
)

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	// Implement admin-related logic here using config.GetConfig().DB
	w.WriteHeader(200)
	w.Write([]byte("admin handler"))
}
