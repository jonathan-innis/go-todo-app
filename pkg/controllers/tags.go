package controllers

import (
	"net/http"

	"github.com/jonathan-innis/go-todo-app/pkg/services"
)

type TagController struct {
	tagService *services.TagService
}

func NewTagController(tagService *services.TagService) *TagController {
	return &TagController{tagService: tagService}
}

func (tc *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
}

func (tc *TagController) GetTags(w http.ResponseWriter, r *http.Request) {
}
