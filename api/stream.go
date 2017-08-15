package api

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
)

const PageSize = 150

func (api *API) GetImageStream(w http.ResponseWriter, r *http.Request) {
	type ResponseEntry struct {
		Url string `json:"url"`
	}

	type ResponseEntryList []ResponseEntry

	type Response struct {
		Data ResponseEntryList `json:"data"`
		PageSize int64 `json:"pageSize"`
		NextMarker *string `json:"nextMarker"`
	}

	// Set marker option for pagination when given.
	marker := r.URL.Query().Get("marker")

	// Do request to S3 instance.
	req := api.s3.ListObjects("stream/", PageSize, marker)

	// Prepare response payload.
	response := Response{
		Data: ResponseEntryList{},
		PageSize: PageSize,
		NextMarker: req.NextMarker,
	}

	for _, obj := range req.Contents {
		key := aws.StringValue(obj.Key)
		if (strings.Contains(key, ".gif")) {
			response.Data = append(response.Data, ResponseEntry{
				Url: api.s3.KeyToUrl(key),
			})
		}
        }

        // Return JSON with signed URL.
	api.WriteJSONResponse(response, w)
}
