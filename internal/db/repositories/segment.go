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
	sql_query := "INSERT INTO segments (slug) VALUES ($1)"
	_, err := sr.db.Exec(sql_query, slug)
	if err != nil {
		return err
	}

	return nil
}

func (sr *SegmentRepository) DeleteSegment(slug string) error {
	sql_query := "DELETE FROM segments WHERE slug = $1"
	_, err := sr.db.Exec(sql_query, slug)
	if err != nil {
		return errors.ErrNotFound
	}

	return nil
}
