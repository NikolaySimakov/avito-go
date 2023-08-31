package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NikolaySimakov/avito-go/internal/db"
	// "github.com/NikolaySimakov/avito-go/internal/models"
)

type SegmentHandler struct {
	repository db.Segment
}

func NewSegmentHandler(r db.Segment) *SegmentHandler {
	return &SegmentHandler{
		repository: r,
	}
}

type createSegmentInput struct {
	Slug string `json:"slug"`
}

type createSegmentResponse struct {
	Slug string `json:"slug"`
}

func (sh *SegmentHandler) Add(w http.ResponseWriter, r *http.Request) {
	var input createSegmentInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	if err := sh.repository.CreateSegment(input.Slug); err != nil {
		panic(err)
	}

	segmentResponse := createSegmentResponse{
		Slug: input.Slug,
	}

	jsonResponse, jsonError := json.Marshal(segmentResponse)

	if jsonError != nil {
		panic(jsonError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type deleteSegmentInput struct {
	Slug string `json:"slug"`
}

func (sh *SegmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var input deleteSegmentInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}

	if err := sh.repository.DeleteSegment(input.Slug); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
