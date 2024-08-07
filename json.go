package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("json marshal failed \n%v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Println("responding with 5xx error code ")
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errorResponse{
		Error: msg,
	})
}
