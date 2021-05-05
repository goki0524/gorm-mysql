package services

import (
	"fmt"

	"github.com/goki0524/gorm-mysql/app/models"
	"github.com/goki0524/gorm-mysql/app/params"
)

type UserService struct {
	service
}

func (s *UserService) GetUser(req *params.GetUserReq) *params.GetUserRes {

	db := s.mysql.Connection()

	var user models.User

	db.First(&user, req.UserID)
	fmt.Println(user)

	var res params.GetUserRes
	res.User = &user

	return &res
}

func (s *UserService) GetUsers() *params.GetUsersRes {

	db := s.mysql.Connection()

	var users []models.User

	db.Where("is_valid = ?", true).Find(&users)
	fmt.Println(users)

	var res params.GetUsersRes
	res.Users = &users

	return &res
}
