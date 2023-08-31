package repositories

import (
	"database/sql"

	"github.com/NikolaySimakov/avito-go/internal/db/errors"
)

type SegmentRepository struct {
	db *sql.DB
}

func NewSegmentRepository(db *sql.DB) *SegmentRepository {
	return &SegmentRepository{
		db: db,
	}
}

func (sr *SegmentRepository) CreateSegment(slug string) error {
	sqlQuery := "INSERT INTO segments (slug) VALUES ($1)"
	_, err := sr.db.Exec(sqlQuery, slug)
	if err != nil {
		return err
	}

	return nil
}

func (sr *SegmentRepository) segmentExist(slug string) (bool, error) {
	sqlQuery := "SELECT slug FROM segments WHERE slug = $1"
	err := sr.db.QueryRow(sqlQuery, slug).Scan(&slug)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, errors.ErrDatabase
		}

		return false, errors.ErrSegmentsNotExist
	}

	return true, nil
}

func (sr *SegmentRepository) DeleteSegment(slug string) error {
	sqlQuery := "DELETE FROM segments WHERE slug = $1"
	_, err := sr.db.Exec(sqlQuery, slug)
	if err != nil {
		return errors.ErrNotFound
	}

	return nil
}
