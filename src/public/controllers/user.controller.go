package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quanndh/go-app/adapter/dtos"
	"github.com/quanndh/go-app/public/services"
	"log"
	"net/http"
)

type UserController struct {
	userService services.IUserService
	logger      *log.Logger
}

func NewUserController(userService services.IUserService, logger *log.Logger) *UserController {
	return &UserController{
		userService: userService,
		logger:      logger,
	}
}

func (c UserController) SignUp(ctx *gin.Context) {
	var body dtos.SignupDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		c.logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid body input",
		})
		return
	}

	user, err1 := c.userService.CreateUser(body)

	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err1,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": user})

}
