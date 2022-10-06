package repositories

import (
	"fmt"
	"github.com/quanndh/go-app/adapter/dtos"
	"github.com/quanndh/go-app/adapter/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (rp UserRepository) CreateUser(data dtos.SignupDto) (*models.User, error) {
	fmt.Println(rp.db)

	user := models.User{Username: data.Username, Password: data.Password}

	res := rp.db.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil

}
