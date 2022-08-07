package main

import (
	"context"
	"flag"
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

	// Define the flags to start the service
	var configType string
	flag.StringVar(&configType, "config-type", string(database.ConfigTypeDefault), "The type of config to use")
	flag.Parse()

	// Get the database config by type
	var config *database.Config
	switch configType {
	case string(database.ConfigTypeEnvironment):
		config = database.NewConfigFromEnvironment()
	default:
		config = database.DefaultConfig()
	}

	db := database.ConnectDB(context.Background(), config)

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
	uc := controllers.NewUserController(us)
	pc := controllers.NewPublicController(is)

	r.Use(middleware.HeaderMiddleware)
	r.Use(middleware.LoggingMiddleware)

	// Frontend/Public Router
	r.HandleFunc("/", pc.Get)

	apiRouter := r.PathPrefix("/api").Subrouter()

	apiV1Router := apiRouter.PathPrefix("/v1").Subrouter()

	apiV1Router.HandleFunc("/login", ac.Login).Methods("POST")
	apiV1Router.HandleFunc("/register", ac.Register).Methods("POST")

	userContextRouter := apiV1Router.PathPrefix("/users/{userId}").Subrouter()
	userContextRouter.Use(middleware.UserAuthenticationMiddleware)

	// Item routes
	userContextRouter.HandleFunc("", uc.GetUser).Methods("GET")
	userContextRouter.HandleFunc("/items/{id}", ic.GetItem).Methods("GET")
	userContextRouter.HandleFunc("/items", ic.ListItems).Methods("GET")
	userContextRouter.HandleFunc("/items", ic.ListItems).Methods("GET").Queries("completed", "{completed}", "listId", "{listId}")
	userContextRouter.HandleFunc("/items", ic.CreateItem).Methods("POST")
	userContextRouter.HandleFunc("/items/{id}", ic.UpdateItem).Methods("PUT")
	userContextRouter.HandleFunc("/items/{id}", ic.DeleteItem).Methods("DELETE")

	// List routes
	userContextRouter.HandleFunc("/lists/{id}", lc.GetList).Methods("GET")
	userContextRouter.HandleFunc("/lists", lc.ListLists).Methods("GET")
	userContextRouter.HandleFunc("/lists", lc.CreateList).Methods("POST")
	userContextRouter.HandleFunc("/lists/{id}", lc.UpdateList).Methods("PUT")
	userContextRouter.HandleFunc("/lists/{id}", lc.DeleteList).Methods("DELETE")

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
