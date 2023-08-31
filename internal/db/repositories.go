package db

import (
	"database/sql"

	"github.com/NikolaySimakov/avito-go/internal/db/repositories"
)

type Segment interface {
	CreateSegment(slug string) error
	DeleteSegment(slug string) error
}

type User interface {
	AddUserSegments(userId string, slugs []string) error
	DeleteUserSegments(userId string, slugs []string) error
	CreateUserIfNotExist(userId string) error
	GetUserSegments(userId string) ([]string, error)
}

type Repositories struct {
	Segment
	User
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Segment: repositories.NewSegmentRepository(db),
		User:    repositories.NewUserRepository(db),
	}
}
