package dto

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID          int64      `json:"id"`
	UUID        uuid.UUID  `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Name        string     `json:"name"`
	Link        string     `json:"link"`
	IsVideo     bool       `json:"is_video"`
	VideoSource string     `json:"video_source"`
	VideoTarget string     `json:"video_target"`
	Editor      string     `json:"editor"`
	PublishedAt *time.Time `json:"published_at"`
	PublishedBy string     `json:"published_by"`
	CreatedAt   *time.Time `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   string     `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   string     `json:"deleted_by"`
	Article     []Article  `json:"articles"`
}
