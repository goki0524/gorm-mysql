package services

import "fmt"

type AuthService struct {
	service
}

func (s *AuthService) Login() {
	fmt.Println("Login!!")
}
