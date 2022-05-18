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

func (p Priority) String() string {
	switch p {
	case LowPriority:
		return "Low"
	case MediumPriority:
		return "Medium"
	case HighPriority:
		return "High"
	default:
		return "None"
	}
}

func ParsePriority(p string) Priority {
	switch p {
	case "Low":
		return LowPriority
	case "Medium":
		return MediumPriority
	case "High":
		return HighPriority
	default:
		return NoPriority
	}
}

type Item struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Title       string               `json:"title" bson:"title" validate:"required"`
	Description string               `json:"description" bson:"description"`
	Priority    Priority             `json:"priority" bson:"priority"`
	DueAt       time.Time            `json:"dueAt" bson:"dueAt"`
	Completed   bool                 `json:"completed" bson:"completed"`
	Tags        []primitive.ObjectID `json:"tags" bson:"tags"`
	ListID      primitive.ObjectID   `json:"listId" bson:"listId"`

	Metadata
}
