package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jonathan-innis/go-todo-app/pkg/constants"
	"net/http"

	"github.com/jonathan-innis/go-todo-app/pkg/auth"
	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"github.com/jonathan-innis/go-todo-app/pkg/services"
	"github.com/jonathan-innis/go-todo-app/pkg/views"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := &views.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
	}

	userId, err := ac.userService.ValidateLoginRequest(context.Background(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		if errors.Is(err, models.UserNotExistsErr{}) || errors.Is(err, models.InvalidPasswordErr{}) {
			helper.GetError(w, http.StatusUnauthorized, "username and password are invalid")
			return
		}
		helper.GetInternalError(w, err)
		return
	}

	tokenStr, err := auth.GetTokenForUserId(userId)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	loginResponse := &views.LoginResponse{
		UserId:    userId,
		Token:     tokenStr,
		TokenType: constants.BearerTokenType,
	}
	_ = json.NewEncoder(w).Encode(loginResponse)
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	registerRequest := &views.RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(registerRequest); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
	}

	if err := helper.ValidateObj(registerRequest); err != nil {
		helper.GetError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err := ac.userService.CreateUser(context.Background(), views.NewUserFromRegisterRequest(registerRequest))
	if errors.Is(err, models.UserExistsErr{}) {
		helper.GetError(w, http.StatusBadRequest, "Username already exists")
		return
	}
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
