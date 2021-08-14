package controller

import (
	"log"
	"net/http"

	"github.com/goki0524/gorm-mysql/app/params"
	"github.com/goki0524/gorm-mysql/app/service"
	"github.com/goki0524/gorm-mysql/app/validator"
	"github.com/goki0524/gorm-mysql/helper/utils"
	"github.com/gorilla/mux"
)

type UserController struct {
	controller
	userService   service.UserService
	userValidator validator.UserValidator
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request, responseを定義
	var req params.GetUserReq
	var res params.GetUserRes

	// パラメーターをrequestに格納
	params := mux.Vars(r)
	strID := params["id"]
	id, err := utils.ConvertStringToInt64(&strID)
	if err != nil {
		log.Println(err)
		return
	}
	req.UserID = id

	// バリデーションを行う
	if ok, err := c.userValidator.GetUser(req); !ok {
		log.Println(err)
		return
	}

	// サービスからresponseを受け取る
	res = *c.userService.GetUser(&req)

	c.controller.ReturnJson(w, &res.User)
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var res params.GetUsersRes
	res = *c.userService.GetUsers()

	c.controller.ReturnJson(w, &res.Users)
}
