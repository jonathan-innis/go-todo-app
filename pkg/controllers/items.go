package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jonathan-innis/go-todo-app/pkg/database"
	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemController struct {
	db database.Interface
}

func NewItemController(db database.Interface) *ItemController {
	return &ItemController{db: db}
}

func (bc *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
		return
	}

	if err := bc.validateItem(item); err != nil {
		helper.GetError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Set the fields that aren't yet set
	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now()
	item.ModifiedAt = time.Now()

	createdID, err := bc.db.Create(context.TODO(), &item)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}

	// Update the ID of the book that we just created
	item.ID = createdID

	w.Header().Add("Location", r.Host+"/api/items/"+item.ID.Hex())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (bc *ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		helper.GetError(w, http.StatusBadRequest, "Invalid request payload with error: "+err.Error())
		return
	}

	if err := bc.validateItem(item); err != nil {
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

		oid, _ := primitive.ObjectIDFromHex(id)

		var oldItem models.Item
		if found, err := bc.db.Get(context.TODO(), id, &oldItem); err != nil {
			helper.GetInternalError(w, err)
			return
		} else if found {
			item.CreatedAt = oldItem.CreatedAt
		} else {
			item.CreatedAt = time.Now()
		}
		item.ID = oid
		item.ModifiedAt = time.Now()

		newCreate, err := bc.db.Update(context.TODO(), item, id)
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

func (bc *ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if id, ok := params["id"]; ok {
		// Validate if the id is a valid ObjectID
		if !primitive.IsValidObjectID(id) {
			helper.GetError(w, http.StatusBadRequest, "ID specified must be a valid ObjectID")
			return
		}

		item := &models.Item{}
		found, err := bc.db.Get(context.TODO(), id, item)
		if err != nil {
			helper.GetInternalError(w, err)
			return
		} else if !found {
			helper.GetError(w, http.StatusNotFound, "Item not found")
			return
		}
		json.NewEncoder(w).Encode(item)
		return
	}
	helper.GetError(w, http.StatusBadRequest, "ID is required")
}

func (bc *ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	items := []models.Item{}
	err := bc.db.List(context.TODO(), &items)
	if err != nil {
		helper.GetInternalError(w, err)
	}
	json.NewEncoder(w).Encode(items)
}

func (bc *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if id, ok := params["id"]; ok {
		if found, err := bc.db.Delete(context.TODO(), id); err != nil {
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

func (bc *ItemController) validateItem(item models.Item) error {
	validate := validator.New()
	return validate.Struct(item)
}
