package dto

type PaginationResponseDto[T any] struct {
	List       []T `json:"list"`
	Limit      int `json:"limit"`
	TotalEntry int `json:"total_entry"`
}

type PaginationResponseDtoExample struct {
	List       []any `json:"list"`
	TotalEntry int   `json:"total_entry" example:"100"`
	Limit      int   `json:"limit" example:"10"`
}
