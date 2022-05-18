package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/services"
	"github.com/jonathan-innis/go-todo-app/pkg/views"
)

type ListController struct {
	listService *services.ListService
}

func NewListController(listService *services.ListService) *ListController {
	return &ListController{listService: listService}
}

func (lc *ListController) CreateList(w http.ResponseWriter, r *http.Request) {
	list := &views.List{}
	if err := json.NewDecoder(r.Body).Decode(list); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
		return
	}

	if err := helper.ValidateObj(list); err != nil {
		helper.GetError(w, http.StatusBadRequest, err.Error())
		return
	}

	listModel, err := lc.listService.CreateItem(context.Background(), views.NewListModel(list))
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	w.Header().Add("Location", r.Host+"/api/items/"+listModel.ID.Hex())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(views.NewListView(listModel))
}

func (lc *ListController) GetLists(w http.ResponseWriter, r *http.Request) {
	lists, err := lc.listService.ListLists(context.Background())
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(views.NewListListView(lists))
}
