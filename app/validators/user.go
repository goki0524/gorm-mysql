package validator

import (
	"github.com/goki0524/gorm-mysql/app/params"
)

type UserValidator struct {
	validator
}

func (v *UserValidator) GetUser(getUserReq params.GetUserReq) (bool, error) {
	validate := v.validator.New()

	if err := validate.Struct(getUserReq); err != nil {
		return false, err
	}

	return true, nil
}
