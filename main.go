package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathan-innis/go-todo-app/pkg/controllers"
	"github.com/jonathan-innis/go-todo-app/pkg/database"
	"github.com/jonathan-innis/go-todo-app/pkg/middleware"
)

func main() {
	r := mux.NewRouter()
	collection := database.ConnectDB(context.Background(), "todo-app", "items")

	db := database.NewMongoDB(collection)
	ic := controllers.NewItemController(db)

	r.Use(middleware.HeaderMiddleware)
	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/api/items", ic.GetBooks).Methods("GET")
	r.HandleFunc("/api/items/{id}", ic.GetBook).Methods("GET")
	r.HandleFunc("/api/items", ic.CreateBook).Methods("POST")
	r.HandleFunc("/api/items/{id}", ic.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/items/{id}", ic.DeleteBook).Methods("DELETE")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))
}
