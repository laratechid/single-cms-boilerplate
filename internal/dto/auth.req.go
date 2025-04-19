package dto

type AuthRequestDto struct {
	Email    string `json:"email" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=5"`
}
