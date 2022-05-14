package services

import (
	"context"
	"time"

	"github.com/jonathan-innis/go-todo-app/pkg/database"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListService struct {
	listCollection database.Interface
}

func NewListService(listCollection database.Interface) *ListService {
	return &ListService{listCollection: listCollection}
}

func (ls *ListService) CreateItem(ctx context.Context, list *models.List) (*models.List, error) {
	// Set the fields that aren't yet set
	list.ID = primitive.NewObjectID()
	list.CreatedAt = time.Now()
	list.ModifiedAt = time.Now()

	createdID, err := ls.listCollection.Create(ctx, &list)
	if err != nil {
		return nil, err
	}
	// Update the ID of the book that we just created
	list.ID = createdID
	return list, nil
}

func (ls *ListService) ListLists(ctx context.Context) ([]models.List, error) {
	lists := []models.List{}
	if err := ls.listCollection.List(ctx, &lists); err != nil {
		return nil, err
	}
	return lists, nil
}
