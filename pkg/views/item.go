package views

import (
	"time"

	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID   `json:"id"`
	Title       string               `json:"title" validate:"required"`
	Description string               `json:"description"`
	Priority    string               `json:"priority"`
	DueAt       time.Time            `json:"dueAt"`
	Completed   bool                 `json:"completed"`
	Tags        []primitive.ObjectID `json:"tags"`

	Metadata
}

func NewItemModel(i *Item) *models.Item {
	return &models.Item{
		ID:          i.ID,
		Title:       i.Title,
		Description: i.Description,
		Priority:    models.ParsePriority(i.Priority),
		DueAt:       i.DueAt,
		Completed:   i.Completed,
		Tags:        i.Tags,
		Metadata:    *NewMetadataModel(&i.Metadata),
	}
}

func NewItemView(i *models.Item) *Item {
	return &Item{
		ID:          i.ID,
		Title:       i.Title,
		Description: i.Description,
		Priority:    i.Priority.String(),
		DueAt:       i.DueAt,
		Completed:   i.Completed,
		Tags:        i.Tags,
		Metadata:    *NewMetadataView(&i.Metadata),
	}
}

func NewItemListView(is []models.Item) []Item {
	var items []Item
	for _, i := range is {
		items = append(items, *NewItemView(&i))
	}
	return items
}
