package v1

import (
	"encoding/json"
	"net/http"

	"github.com/NikolaySimakov/avito-go/internal/db"
	"github.com/gorilla/mux"
	// "github.com/NikolaySimakov/avito-go/internal/models"
)

type SegmentRoutes struct {
	repository db.Segment
}

func NewSegmentRouter(subrouter *mux.Router, r db.Segment) {
	sr := &SegmentRoutes{
		repository: r,
	}

	subrouter.HandleFunc("/", sr.add).Methods("POST")
	subrouter.HandleFunc("/", sr.delete).Methods("DELETE")
}

type createSegmentInput struct {
	Slug string `json:"slug" validate:"required"`
}

type createSegmentOutput struct {
	Slug string `json:"slug"`
}

func (sr *SegmentRoutes) add(w http.ResponseWriter, r *http.Request) {
	var input createSegmentInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	if err := sr.repository.CreateSegment(input.Slug); err != nil {
		panic(err)
	}

	segmentOutput := createSegmentOutput{
		Slug: input.Slug,
	}

	jsonResponse, jsonError := json.Marshal(segmentOutput)

	if jsonError != nil {
		panic(jsonError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type deleteSegmentInput struct {
	Slug string `json:"slug" validate:"required"`
}

func (sr *SegmentRoutes) delete(w http.ResponseWriter, r *http.Request) {
	var input deleteSegmentInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	if err := sr.repository.DeleteSegment(input.Slug); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
