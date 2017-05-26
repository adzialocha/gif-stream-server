package api

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (api *API) GetImageStream(w http.ResponseWriter, r *http.Request) {
	req, _ := api.s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(api.s3BucketName),
	})
	json, _ := json.Marshal(req)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
