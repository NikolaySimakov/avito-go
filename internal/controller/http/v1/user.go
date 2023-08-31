package v1

import (
	"encoding/json"
	"net/http"

	"github.com/NikolaySimakov/avito-go/internal/db"
	"github.com/gorilla/mux"
)

type UserRoutes struct {
	userRepository    db.User
	userTTLRepository db.UserTTL
}

func NewUserRouter(subrouter *mux.Router, r *db.Repositories) {
	ur := &UserRoutes{
		userRepository:    r.User,
		userTTLRepository: r.UserTTL,
	}

	subrouter.HandleFunc("/", ur.show).Methods("GET")
	subrouter.HandleFunc("/", ur.add).Methods("POST")
}

type userSegmentsInput struct {
	UserId         string   `json:"user_id" validate:"required"`
	AddSegments    []string `json:"add_segments" validate:"required"`
	DeleteSegments []string `json:"delete_segments" validate:"required"`
	TTL            uint64   `json:"ttl"`
}

func (ur *UserRoutes) add(w http.ResponseWriter, r *http.Request) {
	var input userSegmentsInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	if err := ur.userRepository.CreateUserIfNotExist(input.UserId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := ur.userRepository.AddUserSegments(input.UserId, input.AddSegments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := ur.userRepository.DeleteUserSegments(input.UserId, input.DeleteSegments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if input.TTL != 0 {
		if err := ur.userTTLRepository.SetTTLForUserSegments(input.UserId, input.AddSegments, int64(input.TTL)); err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}
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
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	// Check TTL
	if err := ur.userTTLRepository.DeleteUserSegments(input.UserId); err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	segments, err := ur.userRepository.GetUserSegments(input.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonResponse, jsonError := json.Marshal(segments)
	if jsonError != nil {
		http.Error(w, jsonError.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
