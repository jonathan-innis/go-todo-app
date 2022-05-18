package views

import (
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username" validate:"required"`
	Email    string             `json:"email" validate:"required"`
	Metadata
}

func NewUserModel(u *User) *models.User {
	return &models.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Metadata: *NewMetadataModel(&u.Metadata),
	}
}

func NewUserView(u *models.User) *User {
	return &User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Metadata: *NewMetadataView(&u.Metadata),
	}
}
