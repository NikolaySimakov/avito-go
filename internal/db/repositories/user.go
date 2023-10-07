package repositories

import (
	"database/sql"

	"github.com/NikolaySimakov/user-segmentation-service/internal/db/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) userExist(userId string) (bool, error) {
	sqlQuery := "SELECT user_id FROM users WHERE user_id = $1"
	err := ur.db.QueryRow(sqlQuery, userId).Scan(&userId)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, errors.ErrDatabase
		}

		return false, nil
	}

	return true, nil
}

func (ur *UserRepository) CreateUserIfNotExist(userId string) error {
	ok, err := ur.userExist(userId)
	if err != nil {
		return err
	}

	if !ok {
		sqlQuery := "INSERT INTO users (user_id) VALUES ($1)"
		_, err := ur.db.Exec(sqlQuery, userId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ur *UserRepository) AddUserSegments(userId string, slugs []string) error {
	sqlQuery := "INSERT INTO user_segments (user_id, segment_slug) VALUES ($1, $2)"

	for _, slug := range slugs {
		_, err := ur.db.Exec(sqlQuery, userId, slug)
		if err != nil {
			return errors.ErrDatabase
		}
	}

	return nil
}

func (ur *UserRepository) DeleteUserSegments(userId string, slugs []string) error {
	sqlQuery := "DELETE FROM user_segments WHERE user_id = $1 AND segment_slug = $2"

	for _, slug := range slugs {
		_, err := ur.db.Exec(sqlQuery, userId, slug)
		if err != nil {
			return errors.ErrDatabase
		}
	}

	return nil
}

func (ur *UserRepository) GetUserSegments(userId string) ([]string, error) {
	ok, err := ur.userExist(userId)

	if !ok {
		if err != nil {
			return []string{}, err
		}

		return []string{}, nil
	}

	sqlQuery := "SELECT segment_slug FROM user_segments WHERE user_id = $1"
	rows, err := ur.db.Query(sqlQuery, userId)
	if err != nil {
		return []string{}, errors.ErrDatabase
	}

	var segments []string = []string{}
	for rows.Next() {
		var slug string
		if err := rows.Scan(&slug); err != nil {
			return []string{}, errors.ErrNotFound
		}

		segments = append(segments, slug)
	}

	return segments, nil
}
