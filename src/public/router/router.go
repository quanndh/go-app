package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quanndh/go-app/public/controllers"
	"github.com/quanndh/go-app/public/middlewares"
	"go.uber.org/fx"
)

type RegisterRouterIn struct {
	fx.In
	Engine *gin.Engine

	UserController *controllers.UserController
}

func RegisterGinRouters(p RegisterRouterIn) {
	group := p.Engine.Group("/")
	group.Use(middlewares.MiddlewareCORS())

	v1 := group.Group("/v1")
	api := v1.Group("/api")

	api.POST("/users", p.UserController.SignUp)

	err := p.Engine.Run("localhost:8000")

	if err != nil {
		panic(err)
	}

}
