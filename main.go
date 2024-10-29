package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Load variables from the .env file
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	databaseURL := envMap["DATABASE_URL"]
	apiKey := envMap["API_KEY"]
	port := envMap["PORT"]

	// Define a handler for the API endpoint
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("Database URL: %s\nAPI Key: %s", databaseURL, apiKey)
		w.Write([]byte(response))
	})

	// Start the server on the specified port
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
