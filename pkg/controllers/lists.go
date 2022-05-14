package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"github.com/jonathan-innis/go-todo-app/pkg/services"
)

type ListController struct {
	listService *services.ListService
}

func NewListController(listService *services.ListService) *ListController {
	return &ListController{listService: listService}
}

func (lc *ListController) CreateList(w http.ResponseWriter, r *http.Request) {
	list := &models.List{}
	if err := json.NewDecoder(r.Body).Decode(list); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
		return
	}

	if err := helper.ValidateObj(list); err != nil {
		helper.GetError(w, http.StatusBadRequest, err.Error())
		return
	}

	item, err := lc.listService.CreateItem(context.Background(), list)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	w.Header().Add("Location", r.Host+"/api/items/"+item.ID.Hex())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (lc *ListController) GetLists(w http.ResponseWriter, r *http.Request) {
	lists, err := lc.listService.ListLists(context.Background())
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(lists)
}
