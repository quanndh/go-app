package boostrap

import (
	"github.com/gin-gonic/gin"
	"github.com/quanndh/go-app/public/controllers"
	"github.com/quanndh/go-app/public/router"
	"go.uber.org/fx"
)

func All() []fx.Option {
	return []fx.Option{

		// Provide controllers, these controllers will be used
		// when register router was invoked
		fx.Provide(controllers.NewUserController),

		// Provide gin engine, register core handlers,
		// actuator endpoints and application routers
		fx.Provide(gin.New),
		fx.Invoke(router.RegisterGinRouters),
	}
}
