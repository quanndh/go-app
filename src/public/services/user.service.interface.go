package services

import (
	"github.com/quanndh/go-app/adapter/dtos"
	"github.com/quanndh/go-app/adapter/models"
)

type IUserService interface {
	CreateUser(data dtos.SignupDto) (*models.User, error)
}
