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

	// Register the collections
	itemCollection := database.NewMongoCollection(db.Collection("items"))
	listCollection := database.NewMongoCollection(db.Collection("lists"))
	userCollection := database.NewMongoCollection(db.Collection("users"))

	// Register the services
	is := services.NewItemService(itemCollection)
	ls := services.NewListService(listCollection)
	us := services.NewUserService(userCollection)

	// Register the controllers
	ic := controllers.NewItemController(is)
	lc := controllers.NewListController(ls)
	ac := controllers.NewAuthController(us)

	r.Use(middleware.HeaderMiddleware)
	r.Use(middleware.LoggingMiddleware)

	apiRouter := r.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/login", ac.Login).Methods("POST")

	userContextRouter := apiRouter.PathPrefix("/users/{userId}").Subrouter()
	userContextRouter.Use(middleware.UserAuthenticationMiddleware)

	// Item routes
	userContextRouter.HandleFunc("/items", ic.GetItems).Methods("GET")
	userContextRouter.HandleFunc("/items", ic.GetItems).Methods("GET").Queries("completed", "{completed}")
	userContextRouter.HandleFunc("/items", ic.CreateItem).Methods("POST")
	userContextRouter.HandleFunc("/items/{id}", ic.UpdateItem).Methods("PUT")
	userContextRouter.HandleFunc("/items/{id}", ic.DeleteItem).Methods("DELETE")

	// List routes
	userContextRouter.HandleFunc("/lists", lc.GetLists).Methods("GET")
	userContextRouter.HandleFunc("/lists", lc.CreateList).Methods("POST")

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
