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
	if found, err := is.itemCollection.Get(ctx, id, &oldItem); err != nil {
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

func (is *ItemService) ListItems(ctx context.Context) ([]models.Item, error) {
	items := []models.Item{}
	if err := is.itemCollection.List(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (is *ItemService) ListItemsByCompleted(ctx context.Context, completed bool) ([]models.Item, error) {
	items := []models.Item{}
	if err := is.itemCollection.ListWithQuery(ctx, &items, map[string]interface{}{"completed": completed}); err != nil {
		return nil, err
	}
	return items, nil
}

func (is *ItemService) DeleteById(ctx context.Context, id string) (bool, error) {
	return is.itemCollection.Delete(ctx, id)
}
