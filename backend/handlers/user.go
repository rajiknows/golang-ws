package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/rajiknows/vedashala/config"
	_ "github.com/rajiknows/vedashala/config"
	"github.com/rajiknows/vedashala/internal/database"
	"github.com/rajiknows/vedashala/utils"
)

type tokens struct {
	accessToken   string
	refresh_token string
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	subrouter := chi.NewRouter()
	subrouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:*"}, // Allow any port on localhost
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow all headers
		ExposedHeaders:   []string{"Set-Cookie", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	subrouter.Get("/", GetUser)
	subrouter.Post("/login", handleLogin)
	subrouter.Post("/register", handleRegister)
	subrouter.Get("/test-cookie", testCookie)
	// ... other user routes

	subrouter.ServeHTTP(w, r)

}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login attempt")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var body LoginBody

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	db := config.GetConfig().DB
	pwd, err := db.GetPassByEmail(r.Context(), body.Email)
	if err != nil {
		utils.RespondWithError(w, 400, "pass does not exist")
		return
	}
	if body.Password != pwd {
		utils.RespondWithError(w, 400, "wrong password")
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24 * 7)

	claims := &Claims{
		Email: body.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatal("error in making token string %v", err)
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenstring,
		Expires:  expirationTime,
		Domain:   "localhost",
		Path:     "/",
		Secure:   false, // Set to true if using HTTPS
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	// http.SetCookie(w, cookie)
	w.Header().Set("Set-Cookie", cookie.String())
	log.Printf("Setting cookie: %+v", cookie)

	utils.RespondWithError(w, 200, "logged in succesfully")
	log.Println("Login successful, cookie set")
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	db := config.GetConfig().DB
	var body RegisterBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.RespondWithError(w, 404, "unsupported data ")
		return
	}

	/// we do not hash it here, instead we hash it client side and give the hashed password to server

	// password, err := utils.HashPassword(body.Password)
	// if err != nil {
	// 	fmt.Print("password hashing failed")
	// }

	user, err := db.GetUserByEmail(r.Context(), body.Email)
	if err != nil {
		user, err := db.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      body.Email,
			Email:     body.Email,
			Password:  body.Password,
		})

		if err != nil {
			utils.RespondWithError(w, 500, "couldnt create a user")
		}
		w.WriteHeader(200)
		utils.RespondWithJson(w, 200, utils.DatabaseUserToUser(user))
		fmt.Println(user)

	}

	utils.RespondWithError(w, 400, fmt.Sprint("%s email already exists", user.Email))
	http.Redirect(w, r, "v1/user/login", 200)

	//_ := config.GetConfig().DB
	// Registration logic here using db

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Load environment variables
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Get the JWT cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "No token provided")
		return
	}

	// Decode the JWT token
	tokenStr := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Check if user exists in the database
	db := config.GetConfig().DB
	user, err := db.GetUserByEmail(r.Context(), claims.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Respond with the user data
	w.WriteHeader(http.StatusOK)
	utils.RespondWithJson(w, http.StatusOK, utils.DatabaseUserToUser(user))
}

func testCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "test-cookie",
		Value:    "test-value",
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	log.Printf("Test cookie set: %+v", cookie)
	w.Write([]byte("Test cookie set"))
}

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
