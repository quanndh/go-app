package controllers

import (
	"fmt"
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	user, err1 := c.userService.CreateUser(body)

	if err1 != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err1.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func (c UserController) Login(ctx *gin.Context) {
	var body dtos.LoginDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	data, err1 := c.userService.Login(body)
	fmt.Println(data, err1)
	if err1 != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err1.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (c UserController) Me(ctx *gin.Context) {

	userId, exist := ctx.Get("UserId")

	if !exist {
		return
	}

	user, err := c.userService.FindById(userId.(uint))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})

}
