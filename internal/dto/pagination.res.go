package dto

type PaginationResponseDto[T any] struct {
	List          []T     `json:"list"`
	TotalEntry    int64   `json:"total_entry"`
	Limit         int64   `json:"limit"`
	TotalPage     float64 `json:"total_page"`
	IsHasNextPage bool    `json:"is_has_next_page"`
}

type PaginationResponseDtoExample struct {
	List          []any   `json:"list"`
	Limit         int64   `json:"limit"`
	TotalEntry    int64   `json:"total_entry" example:"100"`
	TotalPage     float64 `json:"total_page"`
	IsHasNextPage bool    `json:"is_has_next_page" example:"true"`
}
