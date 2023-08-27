package models

import "strings"

type Segments []string

func (s *Segments) ToString() string {
	return strings.Join(*s, ",")
}
