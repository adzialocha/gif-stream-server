package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const IdParamNeededLength = 16

func (api *API) GetSignedUploadURL(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		SignedUrl string `json:"signedUrl"`
		ObjectKey string `json:"objectKey"`
	}

	// Generate name of file with unique id and timestamp.
	id := r.URL.Query().Get("id")
	if (id == "" || len(id) != IdParamNeededLength) {
		api.WriteJSONErrorResponse(
			"Parameter 'id' is missing or too short.",
			http.StatusUnprocessableEntity,
			w,
		)
		return
	}
	now := time.Now().UTC().Format("2006-01-02T15-04-05Z")
	timestamp := strings.Replace(now, "-", "", -1)
	objectKey := "frames/" + id + "_" + timestamp + ".jpg"

	// Get a signed PUT request to upload files to S3.
	req, _ := api.s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(api.s3BucketName),
		Key: aws.String(objectKey),
	})
	url, _ := req.Presign(5 * time.Minute)

	// Return JSON with signed URL.
	response := Response{
		SignedUrl: url,
		ObjectKey: objectKey,
	}
	api.WriteJSONResponse(response, w)
}
