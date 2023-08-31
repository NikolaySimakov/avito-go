package app

import (
	"github.com/NikolaySimakov/avito-go/internal/db"
	"github.com/NikolaySimakov/avito-go/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(repos *db.Repositories) *mux.Router {
	route := mux.NewRouter()

	// Segments
	segmentHandler := handlers.NewSegmentHandler(repos)
	route.HandleFunc("/segment/", segmentHandler.Add).Methods("POST")
	route.HandleFunc("/segment/", segmentHandler.Delete).Methods("DELETE")

	// Users
	userHandler := handlers.NewUserHandler(repos)
	route.HandleFunc("/user/", userHandler.Show).Methods("GET")
	route.HandleFunc("/user/", userHandler.Add).Methods("POST")
	return route
}
