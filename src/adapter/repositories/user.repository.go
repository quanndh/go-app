package repositories

import (
	"github.com/quanndh/go-app/public/dtos"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) CreateUser(data dtos.SignupDto) dtos.SignupDto {
	return data
}
