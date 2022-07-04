package helper

import (
	"encoding/json"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"net/http"
)

func GetInternalError(w http.ResponseWriter, err error) {
	GetError(w, http.StatusInternalServerError, err.Error())
}

func GetError(w http.ResponseWriter, code int, errorMessage string) {
	var response = models.ErrorResponse{
		ErrorMessage: errorMessage,
		StatusCode:   code,
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	_, _ = w.Write(message)
}
