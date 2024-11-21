package main

import (
	"log"
	"net/http"

	"cs-student-platform/backend/internal/database"
	"cs-student-platform/backend/internal/handlers"
	"cs-student-platform/backend/internal/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found", err)
	}

	// Initialize the database
	database.InitDB()
	defer database.DB.Close()

	// Create a new router
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.EnableCORS)

	// Public routes
	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuth)
	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	// Add more protected routes here

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
