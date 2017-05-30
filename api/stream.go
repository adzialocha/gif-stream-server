package api

import "net/http"

const PageSize = 50

func (api *API) GetImageStream(w http.ResponseWriter, r *http.Request) {
	type ResponseEntry struct {
		Key string `json:"key"`
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
		response.Data = append(response.Data, ResponseEntry{
			Key: *obj.Key,
		})
        }

        // Return JSON with signed URL.
	api.WriteJSONResponse(response, w)
}
