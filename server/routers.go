package server

import (
	"net/http"

	"github.com/goki0524/gorm-mysql/app/controllers"
	"github.com/goki0524/gorm-mysql/server/middlewares"
	"github.com/gorilla/mux"
)

const (
	get  = "GET"
	post = "POST"
)

type Router struct {
	Port           string
	authController controllers.AuthController
	userController controllers.UserController
}

func (r *Router) Start() {
	// Init router
	mux := mux.NewRouter()
	mux.Use(middlewares.Logging)

	// Only Login does not require authentication.
	mux.Path("/api/v1/login").Handler(http.HandlerFunc(r.authController.Login))

	// Require Authentication [/api/v1]
	api := mux.PathPrefix("/api/v1").Subrouter()
	api.Use(middlewares.Auth)
	// User
	api.HandleFunc("/user/{id}", r.userController.GetUser).Methods(get)

	// Requires administrator privileges [/api/v1/admin]
	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(middlewares.IsAdmin)
	admin.HandleFunc("/user/{id}", r.userController.GetUser).Methods(get)
	admin.HandleFunc("/users", r.userController.GetUsers).Methods(get)

	// Run Serve
	http.ListenAndServe(":"+r.Port, mux)
}
