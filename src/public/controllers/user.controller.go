package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quanndh/go-app/public/services"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c UserController) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": c.userService.Test(),
	})
}
