package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/services"
	"github.com/jonathan-innis/go-todo-app/pkg/views"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		helper.GetError(w, http.StatusBadRequest, "userId not specified in the request")
		return
	}
	userModel, found, err := uc.userService.GetUserById(context.Background(), userId)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}
	if !found {
		helper.GetError(w, http.StatusNotFound, fmt.Sprintf("User %v not found", userId))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(views.NewUserView(userModel))
}
