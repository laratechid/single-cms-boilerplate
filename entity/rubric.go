package entity

import (
	"time"

	"gorm.io/gorm"
)

type Rubric struct {
	ID         int32           `gorm:"primary_key;column:id" json:"id"`
	Name       string          `json:"name"`
	Alias      string          `json:"alias"`
	IsActive   bool            `json:"is_active"`
	IsMain     bool            `json:"is_main"`
	CreatedAt  *time.Time      `json:"created_at"`
	CreatedBy  string          `json:"created_by"`
	UpdatedAt  *time.Time      `json:"updated_at"`
	UpdatedBy  string          `json:"updated_by"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at"`
	DeletedBy  string          `json:"deleted_by"`
	SubRubrics []SubRubric     `gorm:"foreignKey:rubric_id;references:id"`
}

func (t *Rubric) TableName() string {
	return "rubrics"
}
