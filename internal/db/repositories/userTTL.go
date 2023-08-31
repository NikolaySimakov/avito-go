package repositories

import (
	"database/sql"
	"time"

	"github.com/NikolaySimakov/avito-go/internal/db/errors"
)

type UserTTLRepository struct {
	db *sql.DB
}

func NewUserTTLRepository(db *sql.DB) *UserTTLRepository {
	return &UserTTLRepository{
		db: db,
	}
}

func (utr *UserTTLRepository) SetTTLForUserSegments(userId string, slugs []string, ttl int64) error {
	currentTime := time.Now().UTC()
	deletion_time := currentTime.Add(time.Duration(ttl) * time.Minute)
	sqlQuery := "INSERT INTO user_ttl (user_id, segment_slug, ttl) VALUES ($1, $2, $3)"

	for _, slug := range slugs {
		_, err := utr.db.Exec(sqlQuery, userId, slug, deletion_time)
		if err != nil {
			return errors.ErrDatabase
		}
	}

	return nil
}

func (utr *UserTTLRepository) DeleteUserSegments(userId string) error {
	currentTime := time.Now().UTC()
	sqlQuery := "SELECT segment_slug FROM user_ttl WHERE ttl < $1 AND user_id = $2"
	rows, err := utr.db.Query(sqlQuery, currentTime, userId)
	if err != nil {
		return errors.ErrDatabase
	}

	for rows.Next() {
		var slug string
		if err := rows.Scan(&slug); err != nil {
			return errors.ErrNotFound
		}
		sqlQuery = "DELETE FROM user_segments WHERE user_id = $1 AND segment_slug = $2"
		_, err = utr.db.Exec(sqlQuery, userId, slug)
		if err != nil {
			return errors.ErrDatabase
		}
	}

	// Drop from user_ttl table all this tasks
	sqlQuery = "DELETE FROM user_ttl WHERE ttl < $1 AND user_id = $2"
	_, err = utr.db.Exec(sqlQuery, currentTime, userId)
	if err != nil {
		return errors.ErrDatabase
	}

	return nil
}
