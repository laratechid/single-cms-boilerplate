package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID            int32           `gorm:"primary_key;column:id" json:"id"`
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
	DeletedAt     *gorm.DeletedAt `json:"deleted_at"`
	DeletedBy     string          `json:"deleted_by"`
	Status        string          `json:"status"`
	UnpublishedAt *time.Time      `json:"unpublished_at"`
	UnpublishedBy string          `json:"unpublished_by"`
	Date          *time.Time      `json:"date"`

	Articles      []*Article      `gorm:"many2many:article_groups;foreignKey:id;joinForeignKey:group_id;references:id;joinReferences:article_id" json:"articles,omitempty"`
	GroupsChild   []*Group        `gorm:"foreignKey:parent_id;references:id" json:"groups_child,omitempty"`
	ArticleGroups []*ArticleGroup `json:"article_groups"`
}

func (t *Group) TableName() string {
	return "groups"
}
