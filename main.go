package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rajiknows/vedashala/config"
	"github.com/rajiknows/vedashala/handlers"
	"github.com/rajiknows/vedashala/utils"
)

func main() {
	config.InitConfig()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerConfig(router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("server starting on port", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("error in serving: ", err)
	}
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
}

func routerConfig(router *chi.Mux) {
	v1router := chi.NewRouter()
	v1router.Get("/healthz", handleReadiness)
	v1router.Mount("/user", http.HandlerFunc(handlers.HandleUser))
	v1router.Mount("/student", http.HandlerFunc(handlers.HandleStudent))
	v1router.Mount("/teacher", http.HandlerFunc(handlers.HandleTeacher))
	v1router.Mount("/admin", http.HandlerFunc(handlers.HandleAdmin))
	router.Mount("/v1", v1router)
}
