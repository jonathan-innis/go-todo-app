package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathan-innis/go-todo-app/pkg/controllers"
	"github.com/jonathan-innis/go-todo-app/pkg/database"
	"github.com/jonathan-innis/go-todo-app/pkg/middleware"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	collection := database.ConnectDB(context.Background(), "todo-app", "items")

	db := database.NewMongoDB(collection)
	ic := controllers.NewItemController(db)

	r.Use(middleware.HeaderMiddleware)
	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/api/items", ic.GetItems).Methods("GET").Queries("completed", "{completed}")
	r.HandleFunc("/api/items/{id}", ic.GetItem).Methods("GET")
	r.HandleFunc("/api/items", ic.CreateItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", ic.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", ic.DeleteItem).Methods("DELETE")

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
