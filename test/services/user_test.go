package services_test

import (
	"testing"
	"time"

	"github.com/goki0524/gorm-mysql/app/params"
	"github.com/goki0524/gorm-mysql/app/services"
	"github.com/goki0524/gorm-mysql/test"
	"github.com/stretchr/testify/assert"
)

func init() {
	test.InitTest()
}

// TestGetUserSuccess :	ユーザー取得に成功
func TestGetUserSuccess(t *testing.T) {

	us := services.UserService{}

	userID := int64(1)
	user := *us.GetUser(&params.GetUserReq{
		UserID: userID,
	})

	assert.Equal(t, user.User.ID, userID)
}

// TestGetUserFailure :	ユーザー取得に失敗
func TestGetUserFailure(t *testing.T) {

	us := services.UserService{}

	userID := int64(time.Now().Unix())
	user := *us.GetUser(&params.GetUserReq{
		UserID: userID,
	})

	assert.NotEqual(t, user.User.ID, userID)
}

// TestGetUsersSuccess : 全ユーザーの取得に成功
func TestGetUsersSuccess(t *testing.T) {

	us := services.UserService{}

	res := *us.GetUsers()

	assert.Equal(t, len(*res.Users), 2)
}
