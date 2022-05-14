package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"github.com/jonathan-innis/go-todo-app/pkg/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemController struct {
	itemService *services.ItemService
}

func NewItemController(itemService *services.ItemService) *ItemController {
	return &ItemController{itemService: itemService}
}

func (ic *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
		return
	}

	if err := helper.ValidateObj(item); err != nil {
		helper.GetError(w, http.StatusBadRequest, err.Error())
		return
	}

	item, err := ic.itemService.CreateItem(context.Background(), item)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	w.Header().Add("Location", r.Host+"/api/items/"+item.ID.Hex())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (ic *ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item *models.Item

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
		return
	}

	if err := helper.ValidateObj(item); err != nil {
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

		newCreate, item, err := ic.itemService.UpdateItem(context.Background(), id, item)
		if err != nil {
			helper.GetInternalError(w, err)
			return
		}

		if newCreate {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(item)
		return
	}
	helper.GetError(w, http.StatusBadRequest, "ID is required")
}

func (ic *ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	var completed *bool
	var listId *string
	completedStr := r.FormValue("completed")
	listIdStr := r.FormValue("listId")

	// Get the completed string as a boolean value
	// if it was passed into the parameters
	if completedStr != "" {
		temp, err := strconv.ParseBool(completedStr)
		if err != nil {
			helper.GetError(w, http.StatusBadRequest, "Completed query value must be a boolean")
			return
		}
		completed = &temp
	}
	if listIdStr != "" {
		listId = &listIdStr
	}
	items, err := ic.itemService.ListItems(context.Background(), completed, listId)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (ic *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if id, ok := params["id"]; ok {
		if found, err := ic.itemService.DeleteById(context.TODO(), id); err != nil {
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
