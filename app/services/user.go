package service

import (
	"github.com/goki0524/gorm-mysql/app/params"
	"github.com/goki0524/gorm-mysql/app/repository/mysql"
)

type UserService struct {
	mysqlUserRepository mysql.MysqlUserRepository
}

func (s *UserService) GetUser(req *params.GetUserReq) *params.GetUserRes {

	user := s.mysqlUserRepository.GetByID(req.UserID)

	var res params.GetUserRes
	res.User = user

	return &res
}

// func (s *UserService) GetUsers() *params.GetUsersRes {

// 	db := s.mysql.Connection()

// 	var users []entity.User

// 	db.Where("is_valid = ?", true).Find(&users)
// 	fmt.Println(users)

// 	var res params.GetUsersRes
// 	res.Users = &users

// 	return &res
// }
