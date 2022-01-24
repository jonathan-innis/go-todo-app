package helper

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetInternalError(w http.ResponseWriter, err error) {
	GetError(w, http.StatusInternalServerError, err.Error())
}

func GetError(w http.ResponseWriter, code int, errorMessage string) {
	var response = ErrorResponse{
		ErrorMessage: errorMessage,
		StatusCode:   code,
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
