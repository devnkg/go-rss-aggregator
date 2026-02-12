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

	// Initialize database
	if err := InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	defer CloseDB()

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

	// Feed routes
	mux.HandleFunc("POST /v1/feeds", handleCreateFeed)
	mux.HandleFunc("GET /v1/feeds", handleGetFeeds)
	mux.HandleFunc("GET /v1/feeds/{id}", handleGetFeed)
	mux.HandleFunc("PUT /v1/feeds/{id}", handleUpdateFeed)
	mux.HandleFunc("DELETE /v1/feeds/{id}", handleDeleteFeed)

	// Article routes
	mux.HandleFunc("GET /v1/articles", handleGetArticles)
	mux.HandleFunc("GET /v1/feeds/{feed_id}/articles", handleGetFeedArticles)

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

// Feed Handlers

// handleCreateFeed creates a new RSS feed
func handleCreateFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var feed Feed
	if err := parseJSON(r, &feed); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if feed.Name == "" || feed.URL == "" {
		respondWithError(w, http.StatusBadRequest, "Name and URL are required")
		return
	}

	if result := DB.Create(&feed); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	respondWithJson(w, http.StatusCreated, feed)
}

// handleGetFeeds retrieves all feeds
func handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var feeds []Feed
	if result := DB.Find(&feeds); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch feeds")
		return
	}

	respondWithJson(w, http.StatusOK, feeds)
}

// handleGetFeed retrieves a single feed by ID
func handleGetFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := r.PathValue("id")
	var feed Feed

	if result := DB.First(&feed, "id = ?", id); result.Error != nil {
		respondWithError(w, http.StatusNotFound, "Feed not found")
		return
	}

	respondWithJson(w, http.StatusOK, feed)
}

// handleUpdateFeed updates an existing feed
func handleUpdateFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := r.PathValue("id")
	var feed Feed

	if result := DB.First(&feed, "id = ?", id); result.Error != nil {
		respondWithError(w, http.StatusNotFound, "Feed not found")
		return
	}

	var updateData Feed
	if err := parseJSON(r, &updateData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if updateData.Name != "" {
		feed.Name = updateData.Name
	}
	if updateData.URL != "" {
		feed.URL = updateData.URL
	}

	if result := DB.Save(&feed); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update feed")
		return
	}

	respondWithJson(w, http.StatusOK, feed)
}

// handleDeleteFeed deletes a feed
func handleDeleteFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := r.PathValue("id")

	if result := DB.Delete(&Feed{}, "id = ?", id); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete feed")
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"message": "Feed deleted successfully"})
}

// Article Handlers

// handleGetArticles retrieves all articles
func handleGetArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var articles []Article
	if result := DB.Find(&articles); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch articles")
		return
	}

	respondWithJson(w, http.StatusOK, articles)
}

// handleGetFeedArticles retrieves articles for a specific feed
func handleGetFeedArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	feedID := r.PathValue("feed_id")
	var articles []Article

	if result := DB.Find(&articles, "feed_id = ?", feedID); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch articles")
		return
	}

	respondWithJson(w, http.StatusOK, articles)
}
