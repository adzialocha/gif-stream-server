package api

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (api *API) GetImageStream(w http.ResponseWriter, r *http.Request) {
	req, _ := api.s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(api.s3BucketName),
		Prefix: aws.String("stream/"),
	})
	api.WriteJSONResponse(req, w)
}
