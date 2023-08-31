package models

import "time"

type User struct {
	Id        string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
