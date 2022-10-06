package services

import "log"

type UserService struct {
	logger *log.Logger
}

func NewUserService(logger *log.Logger) *UserService {
	return &UserService{
		logger: logger,
	}
}

func (s UserService) Test() int {
	return 10
}
