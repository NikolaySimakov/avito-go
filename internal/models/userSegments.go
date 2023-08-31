package models

type UserSegments struct {
	UserId      string `db:"user_id"`
	SegmentSlug string `db:"segment_slug"`
}
