package api

import (
	"encoding/json"
	"net/http"

	"github.com/adzialocha/gif-stream-server/s3"
)

type API struct {
	s3 *s3.S3
}

func New(s3Instance *s3.S3) *API {
	return &API{
		s3: s3Instance,
	}
}

func (api *API) WriteJSONResponse(response interface{}, w http.ResponseWriter) {
	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(json)
}

func (api *API) WriteJSONErrorResponse(message string, statusCode int, w http.ResponseWriter) {
	type ErrorResponse struct {
		Message string `json:"message"`
		StatusCode int `json:"statusCode"`
	}

	response := ErrorResponse{
		Message: message,
		StatusCode: statusCode,
	}

	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(json)
}
