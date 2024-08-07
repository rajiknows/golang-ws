package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rajiknows/vedashala/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}

var apiConfig *APIConfig

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found")
	}
	fmt.Println(dbURL)
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("database connection failed %v", err)
	}

	dbQueries := database.New(conn)
	apiConfig = &APIConfig{
		DB: dbQueries,
	}
}

func GetConfig() *APIConfig {
	return apiConfig
}
