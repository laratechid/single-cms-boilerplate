package dto

type ArticleCreateRequestDto struct {
	Title string `json:"title" validate:"required,min=5"`
	Body  string `json:"body" validate:"required,min=5"`
}

type ArticleUpdateRequestDto struct {
	Title string `json:"title,omitempty" validate:"required,min=5"`
	Body  string `json:"body,omitempty" validate:"required,min=5"`
}
