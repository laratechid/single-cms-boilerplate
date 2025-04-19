package dto

type PaginationRequestDto struct {
	Page  int `json:"page" form:"page" example:"1"`
	Limit int `json:"limit" form:"limit" example:"10"`
}

func (p *PaginationRequestDto) SetDefault() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
}
