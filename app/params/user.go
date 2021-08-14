package params

import (
	"github.com/goki0524/gorm-mysql/app/entity"
)

/* Follow the rules below for validation files. */
/*
	Request:
		function: type ServiceNameReq struct {}
		validate: be sure to do it

	Response:
		function: type ServiceNameRes struct {}
		validate: you don't have to check
*/

// GetUser
type GetUserReq struct {
	UserID int64 `validate:"required"`
}

type GetUserRes struct {
	User *entity.User
}

// GetUsers
type GetUsersReq struct {
}

type GetUsersRes struct {
	Users *[]entity.User
}
