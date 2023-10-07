package app

import (
	v1 "github.com/NikolaySimakov/user-segmentation-service/internal/controller/http/v1"
	"github.com/NikolaySimakov/user-segmentation-service/internal/db"
	"github.com/gorilla/mux"
)

func NewRouter(repos *db.Repositories) *mux.Router {
	route := mux.NewRouter()

	// Segments subrouter
	segmentSubrouter := route.PathPrefix("/segment").Subrouter()
	v1.NewSegmentRouter(segmentSubrouter, repos)

	// Users subrouter
	userSubrouter := route.PathPrefix("/user").Subrouter()
	v1.NewUserRouter(userSubrouter, repos)

	return route
}
