package entity

import "time"

type Article struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
