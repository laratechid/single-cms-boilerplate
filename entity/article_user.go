package entity

import (
	"time"

	"gorm.io/gorm"
)

type ArticleUser struct {
	ID        int32           `json:"id"`
	ArticleID int32           `json:"article_id"`
	UserID    int32           `json:"user_id"`
	Type      string          `json:"type"`
	CreatedAt *time.Time      `json:"created_at"`
	CreatedBy string          `json:"created_by"`
	UpdatedAt *time.Time      `json:"updated_at"`
	UpdatedBy string          `json:"updated_by"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	DeletedBy string          `json:"deleted_by"`
	User      *User           `json:"user"`
	Article   *Article        `json:"article"`
}

func (t *ArticleUser) TableName() string {
	return "article_user"
}
