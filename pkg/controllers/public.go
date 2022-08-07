package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jonathan-innis/go-todo-app/pkg/helper"
	"github.com/jonathan-innis/go-todo-app/pkg/services"
	"github.com/jonathan-innis/go-todo-app/pkg/views"
)

type PublicController struct {
	itemService *services.ItemService
}

func NewPublicController(itemService *services.ItemService) *PublicController {
	return &PublicController{itemService: itemService}
}

func (pc *PublicController) Get(w http.ResponseWriter, r *http.Request) {
	itemModels, err := pc.itemService.ListItems(context.Background(), nil, nil)
	if err != nil {
		helper.GetInternalError(w, err)
		return
	}
	l := views.NewItemListView(itemModels)

	res := "<html><h1>Todo List</h1><hr>"
	for _, elem := range l {
		res += pc.itemTemplate(elem)
	}
	res += "</html>"
	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(res))
}

func (pc *PublicController) itemTemplate(items views.Item) string {
	return fmt.Sprintf("<h3>%s: %s</h3>", items.Title, items.Description)
}
