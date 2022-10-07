package services

import (
	"github.com/quanndh/go-app/adapter/dtos"
	"github.com/quanndh/go-app/public/resources"
)

type IUserService interface {
	CreateUser(data dtos.SignupDto) (*resources.UserResource, error)
	Login(data dtos.LoginDto) (*resources.LoginResource, error)
}
