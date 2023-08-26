package app

import (
	"github.com/NikolaySimakov/avito-go/internal/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", handlers.AddUser).Methods("POST")
	route.HandleFunc("/{id}", handlers.GetUserSegments).Methods("GET")
	route.HandleFunc("/{id}", handlers.EditUserSegments).Methods("PUT")
	route.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")
	return route
}
