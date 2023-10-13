package helper

import (
	"encoding/json"
	"net/http"
)

func ToResponseBody(writer http.ResponseWriter, response interface{}, statusCode int) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
