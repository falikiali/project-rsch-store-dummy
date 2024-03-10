package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, requestBody interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
