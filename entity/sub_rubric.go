package entity

import (
	"time"

	"gorm.io/gorm"
)

type SubRubric struct {
	ID        int32           `gorm:"primary_key,column:id" json:"id"`
	Name      string          `json:"name"`
	Alias     string          `json:"alias"`
	IsActive  bool            `json:"is_active"`
	RubricID  int32           `json:"rubric_id"`
	CreatedAt *time.Time      `json:"created_at"`
	CreatedBy string          `json:"created_by"`
	UpdatedAt *time.Time      `json:"updated_at"`
	UpdatedBy string          `json:"updated_by"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	DeletedBy string          `json:"deleted_by"`
	Articles  []Article       `gorm:"foreignKey:sub_rubric_id;references:id"`
	Rubric    Rubric          `gorm:"foreignKey:RubricID;references:id"`
}

func (t *SubRubric) TableName() string {
	return "sub_rubrics"
}
