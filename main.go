package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathan-innis/go-todo-app/pkg/controllers"
	"github.com/jonathan-innis/go-todo-app/pkg/database"
	"github.com/jonathan-innis/go-todo-app/pkg/middleware"
	"github.com/jonathan-innis/go-todo-app/pkg/services"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	db := database.ConnectDB(context.Background(), "todo-app")

	itemCollection := database.NewMongoCollection(db.Collection("items"))
	tagCollection := database.NewMongoCollection(db.Collection("tags"))

	is := services.NewItemService(itemCollection)
	ic := controllers.NewItemController(is)

	ts := services.NewTagService(tagCollection)
	tc := controllers.NewTagController(ts)

	r.Use(middleware.HeaderMiddleware)
	r.Use(middleware.LoggingMiddleware)

	// Item routes
	r.HandleFunc("/api/items", ic.GetItems).Methods("GET").Queries("completed", "{completed}")
	r.HandleFunc("/api/items", ic.CreateItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", ic.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", ic.DeleteItem).Methods("DELETE")

	// Tag routes
	r.HandleFunc("/api/tags", tc.GetTags).Methods("GET")
	r.HandleFunc("/api/tags", tc.CreateTag).Methods("POST")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{"*"},
	})
	// set our port address
	log.Fatal(http.ListenAndServe(":8000", corsOpts.Handler(r)))
}
