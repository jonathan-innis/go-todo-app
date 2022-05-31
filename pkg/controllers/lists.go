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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (lc *ListController) UpdateList(w http.ResponseWriter, r *http.Request) {
	var list *views.List

	if err := json.NewDecoder(r.Body).Decode(list); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
		return
	}

	if err := helper.ValidateObj(list); err != nil {
		helper.GetError(w, http.StatusBadRequest, err.Error())
		return
	}

	params := mux.Vars(r)
	if id, ok := params["id"]; ok {
		// Validate if the id is a valid ObjectID
		if !primitive.IsValidObjectID(id) {
			helper.GetError(w, http.StatusBadRequest, "ID specified must be a valid ObjectID")
			return
		}

		newCreate, listModel, err := lc.listService.UpdateList(context.Background(), id, views.NewListModel(list))
		if err != nil {
			helper.GetInternalError(w, err)
			return
		}

		if newCreate {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(views.NewListView(listModel))
		return
	}
	helper.GetError(w, http.StatusBadRequest, "ID is required")
}

func (lc *ListController) GetList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		helper.GetError(w, http.StatusBadRequest, "id not specified in the request")
		return
	}
	listModel, found, err := lc.listService.GetListById(context.Background(), id)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}
	if !found {
		helper.GetError(w, http.StatusNotFound, fmt.Sprintf("Item %v not found", id))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(views.NewListView(listModel))
}

func (lc *ListController) ListLists(w http.ResponseWriter, r *http.Request) {
	lists, err := lc.listService.ListLists(context.Background())
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(views.NewListListView(lists))
}

func (lc *ListController) DeleteList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if id, ok := params["id"]; ok {
		if found, err := lc.listService.DeleteById(context.TODO(), id); err != nil {
			helper.GetInternalError(w, err)
			return
		} else if !found {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	helper.GetError(w, http.StatusBadRequest, "ID is required")
}
