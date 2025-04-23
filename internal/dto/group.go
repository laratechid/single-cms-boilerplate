package dto

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID            int32           `json:"id"`
	UUID          uuid.UUID       `json:"uuid"`
	Level         int32           `json:"level"`
	Cover         string          `json:"cover"`
	Caption       string          `json:"caption"`
	Category      string          `json:"category"`
	Alias         string          `json:"alias"`
	Description   string          `json:"description"`
	MetaTitle     string          `json:"meta_title"`
	MetaKeyword   string          `json:"meta_keyword"`
	Sequence      int32           `json:"sequence"`
	IsActive      bool            `json:"is_active"`
	PublishedAt   *time.Time      `json:"published_at"`
	PublishedBy   string          `json:"published_by"`
	ParentID      int32           `json:"parent_id"`
	CreatedAt     *time.Time      `json:"created_at"`
	CreatedBy     string          `json:"created_by"`
	UpdatedAt     *time.Time      `json:"updated_at"`
	UpdatedBy     string          `json:"updated_by"`
	DeletedAt     *time.Time      `json:"deleted_at"`
	DeletedBy     string          `json:"deleted_by"`
	Status        string          `json:"status"`
	UnpublishedAt *time.Time      `json:"unpublished_at"`
	UnpublishedBy string          `json:"unpublished_by"`
	Date          *time.Time      `json:"date"`
	Articles      []*Article      `json:"articles"`
	GroupsChild   []*Group        `json:"groups_child"`
	ArticleGroups []*ArticleGroup `json:"article_groups"`
}
