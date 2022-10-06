package services

import (
	"github.com/quanndh/go-app/adapter/dtos"
	"github.com/quanndh/go-app/adapter/models"
	"github.com/quanndh/go-app/adapter/repositories"
	"log"
)

type UserService struct {
	logger         *log.Logger
	userRepository repositories.UserRepository
}

func NewUserService(logger *log.Logger) IUserService {
	return &UserService{
		logger: logger,
	}
}

func (s UserService) CreateUser(data dtos.SignupDto) (*models.User, error) {
	user, err := s.userRepository.CreateUser(data)

	if err != nil {
		return nil, err
	}

	return user, err
}
