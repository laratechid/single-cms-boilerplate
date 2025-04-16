package dto

type PaginationResponseDto[T any] struct {
	Data       []T `json:"data"`
	Limit      int `json:"limit"`
	TotalEntry int `json:"total_entry"`
}
