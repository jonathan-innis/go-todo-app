package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Priority int64

const (
	NoPriority Priority = iota
	LowPriority
	MediumPriority
	HighPriority
)

type Item struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Title       string               `json:"title" bson:"title" validate:"required"`
	Description string               `json:"description" bson:"description"`
	Priority    Priority             `json:"priority" bson:"priority"`
	DueAt       time.Time            `json:"dueAt" bson:"dueAt"`
	Completed   bool                 `json:"completed" bson:"completed"`
	Tags        []primitive.ObjectID `json:"tags" bson:"tags"`
	CreatedAt   time.Time            `json:"createdAt" bson:"createdAt"`
	ModifiedAt  time.Time            `json:"modifiedAt" bson:"modifiedAt"`
}
