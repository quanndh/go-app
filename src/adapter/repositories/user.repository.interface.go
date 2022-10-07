package repositories

import (
	"github.com/quanndh/go-app/adapter/dtos"
	"github.com/quanndh/go-app/adapter/models"
)

type IUserRepository interface {
	CreateUser(data dtos.SignupDto) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}
