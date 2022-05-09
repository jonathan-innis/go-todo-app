package services

import "github.com/jonathan-innis/go-todo-app/pkg/database"

type TagService struct {
	tagCollection database.Interface
}

func NewTagService(tagCollection database.Interface) *TagService {
	return &TagService{tagCollection: tagCollection}
}
