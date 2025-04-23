package dto

import "time"

type Rubric struct {
	ID         int32       `json:"id"`
	Name       string      `json:"name"`
	Alias      string      `json:"alias"`
	IsActive   bool        `json:"is_active"`
	IsMain     bool        `json:"is_main"`
	CreatedAt  *time.Time  `json:"created_at"`
	CreatedBy  string      `json:"created_by"`
	UpdatedAt  *time.Time  `json:"updated_at"`
	UpdatedBy  string      `json:"updated_by"`
	DeletedAt  *time.Time  `json:"deleted_at"`
	DeletedBy  string      `json:"deleted_by"`
	SubRubrics []SubRubric `json:"sub_rubric"`
}

type SubRubric struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name"`
	Alias     string     `json:"alias"`
	IsActive  bool       `json:"is_active"`
	RubricID  int32      `json:"rubric_id"`
	CreatedAt *time.Time `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy string     `json:"deleted_by"`
	Articles  []Article  `json:"articles"`
	Rubric    Rubric     `json:"rubric"`
}
