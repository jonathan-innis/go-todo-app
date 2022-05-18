package controllers

import (
	"context"
	"encoding/json"
	"errors"
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

	// TODO: grab the user from the DB with the user service and validate that the password matches

	userId := "xxxxx"
	tokenStr, err := auth.GetTokenForUserId(userId)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	loginResponse := &views.LoginResponse{
		UserId: userId,
		Token:  tokenStr,
	}
	json.NewEncoder(w).Encode(loginResponse)
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	user := &views.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
	}

	if err := helper.ValidateObj(user); err != nil {
		helper.GetError(w, http.StatusBadRequest, err.Error())
		return
	}

	userModel, err := ac.userService.CreateUser(context.Background(), views.NewUserModel(user))
	if errors.Is(err, models.UserExistsErr{}) {
		helper.GetError(w, http.StatusBadRequest, "Username already exists")
		return
	}
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(views.NewUserView(userModel))
}
