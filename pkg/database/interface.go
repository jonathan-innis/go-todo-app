package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interface interface {
	Create(ctx context.Context, item interface{}) (primitive.ObjectID, error)
	Update(ctx context.Context, item interface{}, idStr string) (bool, error)
	List(ctx context.Context, items interface{}) error
	ListWithQuery(ctx context.Context, items interface{}, query map[string]interface{}) error
	Get(ctx context.Context, id string, item interface{}) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}
