package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathan-innis/go-todo-app/pkg/auth"
	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := &models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
	}

	userId := "xxxxx"
	tokenStr, err := auth.GetTokenForUserId(userId)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	loginResponse := &models.LoginResponse{
		UserId: userId,
		Token:  tokenStr,
	}
	json.NewEncoder(w).Encode(loginResponse)
}
