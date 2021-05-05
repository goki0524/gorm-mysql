package controllers

import (
	"net/http"

	"github.com/goki0524/gorm-mysql/app/services"
)

type AuthController struct {
	controller
	authService services.AuthService
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	c.authService.Login()
}
