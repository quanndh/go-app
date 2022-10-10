package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quanndh/go-app/public/config"
	"github.com/quanndh/go-app/public/controllers"
	"github.com/quanndh/go-app/public/middlewares"
	"github.com/quanndh/go-app/public/services"
	"go.uber.org/fx"
)

type RegisterRouterIn struct {
	fx.In
	Engine *gin.Engine

	JwtService     services.IJwtService
	UserController *controllers.UserController
	Config         *config.Configuration
}

func RegisterGinRouters(p RegisterRouterIn) {
	group := p.Engine.Group("/")
	group.Use(middlewares.MiddlewareCORS())

	v1 := group.Group("/v1")
	api := v1.Group("/api")

	api.POST("/login", p.UserController.Login)
	api.POST("/users", p.UserController.SignUp)
	api.GET("/me", middlewares.MiddlewareAuthentication(p.JwtService), p.UserController.Me)

	err := p.Engine.Run("localhost:8000")

	if err != nil {
		panic(err)
	}

}
