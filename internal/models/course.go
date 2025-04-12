package models

import "time"

type Course struct {
	Id          string
	Title       string
	Description string
	Creator_id  string
	Created_at  time.Time
}
