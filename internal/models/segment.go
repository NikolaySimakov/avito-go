package models

import "time"

type Segment struct {
	Slug      string    `db:"slug"`
	CreatedAt time.Time `db:"created_at"`
}
