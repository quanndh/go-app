package repositories

import (
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

	user := models.User{Username: data.Username, Password: data.Password}

	res := rp.db.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (rp UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	res := rp.db.First(&user, "username = ?", username)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (rp UserRepository) FindById(id uint) (*models.User, error) {
	var user models.User

	res := rp.db.First(&user, id)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
