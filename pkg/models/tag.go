package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	ModifiedAt  time.Time          `json:"modifiedAt" bson:"modifiedAt"`
}
