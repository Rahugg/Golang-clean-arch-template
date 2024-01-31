package user

import (
	"github.com/Rahugg/Golang-clean-arch-template/internal/repository"
	def "github.com/Rahugg/Golang-clean-arch-template/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}