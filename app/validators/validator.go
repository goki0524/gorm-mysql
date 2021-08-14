package validator

import (
	playgroundValidator "github.com/go-playground/validator"
)

type validator struct{}

func (v *validator) New() *playgroundValidator.Validate {

	validate := playgroundValidator.New()
	return validate
}
