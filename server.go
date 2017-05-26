package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

type API struct {
	s3Client *s3.S3
}

func (api *API) upload(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		SignedUrl string `json:"signedUrl"`
		ObjectKey string `json:"objectKey"`
	}

	// @TODO
	objectKey := "test"

	// Get a signed PUT request to upload files to S3.
	req, _ := api.s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(objectKey),
	})
	url, _ := req.Presign(15 * time.Minute)

	// Return JSON with signed URL.
	response := Response{
		SignedUrl: url,
		ObjectKey: objectKey,
	}
	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (api *API) stream(w http.ResponseWriter, r *http.Request) {
	req, _ := api.s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
	})
	json, _ := json.Marshal(req)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	// Load environment variables
	godotenv.Load()

	// Set port of HTTP server.
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Prepare S3 configuration.
	creds := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("AWS_SESSION_TOKEN"),
	)
	cfg := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")).WithCredentials(creds)

	// Create API with S3 client.
	api := &API{
		s3Client: s3.New(session.New(), cfg),
	}

	// Start HTTP static file server.
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/upload", api.upload)
	http.HandleFunc("/api/stream", api.stream)
	http.ListenAndServe(":"+port, nil)
}
