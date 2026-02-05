package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load(".env")

	// Get the port from environment or use default
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
		log.Println("PORT not found, using default: 8080")
	}

	// Create a new mux (router) for handling different routes
	mux := http.NewServeMux()

	// Register route handlers
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)

	// Start the server
	log.Printf("Server starting on port %s\n", portString)
	log.Fatal(http.ListenAndServe(":"+portString, mux))
}

// handleRoot handles requests to the root path
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to RSS Aggregator API\n")
}

// handleHealth handles requests to the /health endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
