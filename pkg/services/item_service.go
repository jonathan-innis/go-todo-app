package services

import (
	"context"
	"time"

	"github.com/jonathan-innis/go-todo-app/pkg/database"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemService struct {
	itemCollection database.Interface
}

func NewItemService(itemCollection database.Interface) *ItemService {
	return &ItemService{itemCollection: itemCollection}
}

func (is *ItemService) CreateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	// Set the fields that aren't yet set
	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now()
	item.ModifiedAt = time.Now()

	createdID, err := is.itemCollection.Create(ctx, &item)
	if err != nil {
		return nil, err
	}
	// Update the ID of the book that we just created
	item.ID = createdID
	return item, nil
}

func (is *ItemService) UpdateItem(ctx context.Context, id string, item *models.Item) (bool, *models.Item, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var oldItem models.Item
	if found, err := is.itemCollection.GetById(ctx, id, &oldItem); err != nil {
		return false, nil, err
	} else if found {
		item.CreatedAt = oldItem.CreatedAt
	} else {
		item.CreatedAt = time.Now()
	}
	item.ID = oid
	item.ModifiedAt = time.Now()

	newCreate, err := is.itemCollection.Update(ctx, item, id)
	if err != nil {
		return false, nil, err
	}
	return newCreate, item, nil
}

func (is *ItemService) GetItemById(ctx context.Context, id string) (*models.Item, bool, error) {
	item := &models.Item{}
	found, err := is.itemCollection.GetById(ctx, id, item)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}
	return item, true, nil
}

func (is *ItemService) ListItems(ctx context.Context, completed *bool, listId *string) ([]models.Item, error) {
	items := []models.Item{}
	queryParams, err := is.getQueryParams(completed, listId)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		if err := is.itemCollection.ListWithQuery(ctx, &items, queryParams); err != nil {
			return nil, err
		}
	} else {
		if err := is.itemCollection.List(ctx, &items); err != nil {
			return nil, err
		}
	}
	return items, nil
}

func (is *ItemService) getQueryParams(completed *bool, listId *string) (map[string]interface{}, error) {
	queries := make(map[string]interface{})

	if completed != nil {
		queries["completed"] = *completed
	}
	if listId != nil {
		listOId, err := primitive.ObjectIDFromHex(*listId)
		if err != nil {
			return nil, err
		}
		queries["listId"] = listOId
	}
	return queries, nil
}

func (is *ItemService) DeleteById(ctx context.Context, id string) (bool, error) {
	return is.itemCollection.Delete(ctx, id)
}
