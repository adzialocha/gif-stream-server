package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (api *API) GetSignedUploadURL(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		SignedUrl string `json:"signedUrl"`
		ObjectKey string `json:"objectKey"`
	}

	// @TODO
	objectKey := "test"

	// Get a signed PUT request to upload files to S3.
	req, _ := api.s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(api.s3BucketName),
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
