package services

import (
	"context"
	"time"

	"github.com/jonathan-innis/go-todo-app/pkg/database"
	"github.com/jonathan-innis/go-todo-app/pkg/errors"
	"github.com/jonathan-innis/go-todo-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userCollection database.Interface
}

func NewUserService(userCollection database.Interface) *UserService {
	return &UserService{userCollection: userCollection}
}

func (us *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, found, err := us.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	if found {
		return nil, errors.UserExistsErr{}
	}

	// Set the fields that aren't yet set
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()

	createdID, err := us.userCollection.Create(ctx, &user)
	if err != nil {
		return nil, err
	}
	// Update the ID of the user that we just created
	user.ID = createdID
	return user, nil
}

func (us *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, bool, error) {
	user := &models.User{}
	found, err := us.userCollection.GetOneWithQuery(ctx, map[string]interface{}{"username": username}, user)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}
	return user, true, nil
}
