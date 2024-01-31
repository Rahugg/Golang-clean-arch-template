package app

import (
	"github.com/Rahugg/Golang-clean-arch-template/internal/api/user"
	"github.com/Rahugg/Golang-clean-arch-template/internal/repository"
	userRepository "github.com/Rahugg/Golang-clean-arch-template/internal/repository/user"
	"github.com/Rahugg/Golang-clean-arch-template/internal/service"
	userService "github.com/Rahugg/Golang-clean-arch-template/internal/service/user"
	"github.com/jmoiron/sqlx"
)

type serviceProvider struct {
	userRepository repository.UserRepository
	userService    service.UserService
	userImpl       *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(sqlx.NewDb(nil, ""))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService())
	}

	return s.userImpl
}
