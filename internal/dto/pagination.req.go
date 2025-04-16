package dto

type PaginationRequestDto struct {
	Page  int `json:"page" form:"page" example:"1"`
	Limit int `json:"limit" form:"limit" example:"10"`
}
