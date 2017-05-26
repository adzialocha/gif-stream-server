package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adzialocha/gif-stream-server/api"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables.
	godotenv.Load()
}

func main() {
	// Set port of HTTP server.
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Create API with S3 client.
	apiHandler := api.New(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("AWS_SESSION_TOKEN"),
		os.Getenv("AWS_BUCKET_NAME"),
		os.Getenv("AWS_REGION"),
	)

	// Start HTTP static file server.
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/upload", apiHandler.GetSignedUploadURL)
	http.HandleFunc("/api/stream", apiHandler.GetImageStream)
	http.ListenAndServe(":"+port, nil)
}
