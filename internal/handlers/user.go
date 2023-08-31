package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NikolaySimakov/avito-go/internal/db"
)

type UserHandler struct {
	repository db.User
}

func NewUserHandler(r db.User) *UserHandler {
	return &UserHandler{
		repository: r,
	}
}

type userSegmentsInput struct {
	UserId         string   `json:"user_id"`
	AddSegments    []string `json:"add_segments"`
	DeleteSegments []string `json:"delete_segments"`
}

func (uh *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var input userSegmentsInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	if err := uh.repository.CreateUserIfNotExist(input.UserId); err != nil {
		panic(err)
	}

	if err := uh.repository.AddUserSegments(input.UserId, input.AddSegments); err != nil {
		panic(err)
	}

	if err := uh.repository.DeleteUserSegments(input.UserId, input.DeleteSegments); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

type showSegmentsInput struct {
	UserId string `json:"user_id"`
}

type showSegmentsOutput struct {
	Segments []string `json:"segments"`
}

func (uh *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	var input showSegmentsInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	segments, err := uh.repository.GetUserSegments(input.UserId)
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
