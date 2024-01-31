package user

import "github.com/Rahugg/Golang-clean-arch-template/internal/service"

type Implementation struct {
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
