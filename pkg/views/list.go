package views

import (
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description"`
	Metadata
}

func NewListModel(l *List) *models.List {
	return &models.List{
		ID:          l.ID,
		Name:        l.Name,
		Description: l.Description,
		Metadata:    *NewMetadataModel(&l.Metadata),
	}
}

func NewListView(l *models.List) *List {
	return &List{
		ID:          l.ID,
		Name:        l.Name,
		Description: l.Description,
		Metadata:    *NewMetadataView(&l.Metadata),
	}
}

func NewListListView(ls []models.List) []List {
	var lists []List
	for _, l := range ls {
		lists = append(lists, *NewListView(&l))
	}
	return lists
}
