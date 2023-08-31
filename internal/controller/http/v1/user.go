package v1

import (
	"encoding/json"
	"net/http"

	"github.com/NikolaySimakov/avito-go/internal/db"
	"github.com/gorilla/mux"
)

type UserRoutes struct {
	repository db.User
}

func NewUserRouter(subrouter *mux.Router, r db.User) {
	ur := &UserRoutes{
		repository: r,
	}

	subrouter.HandleFunc("/", ur.show).Methods("GET")
	subrouter.HandleFunc("/", ur.add).Methods("POST")
}

type userSegmentsInput struct {
	UserId         string   `json:"user_id" validate:"required"`
	AddSegments    []string `json:"add_segments" validate:"required"`
	DeleteSegments []string `json:"delete_segments" validate:"required"`
}

func (ur *UserRoutes) add(w http.ResponseWriter, r *http.Request) {
	var input userSegmentsInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	if err := ur.repository.CreateUserIfNotExist(input.UserId); err != nil {
		panic(err)
	}

	if err := ur.repository.AddUserSegments(input.UserId, input.AddSegments); err != nil {
		panic(err)
	}

	if err := ur.repository.DeleteUserSegments(input.UserId, input.DeleteSegments); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

type showSegmentsInput struct {
	UserId string `json:"user_id" validate:"required"`
}

type showSegmentsOutput struct {
	Segments []string `json:"segments"`
}

func (ur *UserRoutes) show(w http.ResponseWriter, r *http.Request) {
	var input showSegmentsInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	segments, err := ur.repository.GetUserSegments(input.UserId)
	if err != nil {
		panic(err)
	}

	jsonResponse, jsonError := json.Marshal(segments)
	if jsonError != nil {
		panic(jsonError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
