package helper

import (
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(payload any) error {
	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(payload); err != nil {
		return err
	}
	return nil
}
