package db

import (
	"database/sql"

	"github.com/NikolaySimakov/avito-go/internal/db/repositories"
)

type Segment interface {
	CreateSegment(slug string) error
	DeleteSegment(slug string) error
}

type Repositories struct {
	Segment
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Segment: repositories.NewSegmentRepository(db),
	}
}
