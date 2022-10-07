package resources

import "github.com/quanndh/go-app/adapter/models"

type UserResource struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func NewUserResource(model *models.User) *UserResource {
	return &UserResource{
		ID:        model.ID,
		Username:  model.Username,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

type LoginResource struct {
	Token string        `json:"token"`
	User  *UserResource `json:"user"`
}

func NewLoginResource(model *models.User, token string) *LoginResource {
	return &LoginResource{
		User:  NewUserResource(model),
		Token: token,
	}
}
