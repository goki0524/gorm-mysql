package mysql

import (
	"fmt"

	"github.com/goki0524/gorm-mysql/app/entity"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

type mysqlUserRepositoryInter interface {
	GetByID(id int64) *entity.User
}

func NewMysqlUserRepository(db *gorm.DB) mysqlUserRepositoryInter {
	return &mysqlUserRepository{db}
}

func (m *mysqlUserRepository) GetByID(id int64) *entity.User {
	var user entity.User
	m.db.First(&user, id)
	fmt.Println(user)

	return &user
}
