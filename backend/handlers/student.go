package handlers

import (
	"net/http"

	_ "github.com/rajiknows/vedashala/config"
)

func HandleStudent(w http.ResponseWriter, r *http.Request) {
	// Implement student-related logic here using config.GetConfig().DB
	w.WriteHeader(200)
	w.Write([]byte("student handler"))
}
