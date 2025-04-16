package dto

type AuthRequestDto struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=5"`
}
