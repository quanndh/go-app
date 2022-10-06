package boostrap

import (
	"github.com/gin-gonic/gin"
	"github.com/quanndh/go-app/public/controllers"
	"github.com/quanndh/go-app/public/router"
	"github.com/quanndh/go-app/public/services"
	"go.uber.org/fx"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "[Log] " /* prefix */, 19 /* flags */)
	return logger
}

func All() []fx.Option {
	return []fx.Option{
		fx.Provide(NewLogger),

		// Provide services
		fx.Provide(services.NewUserService),

		// Provide controllers, these controllers will be used
		// when register router was invoked
		fx.Provide(controllers.NewUserController),

		// Provide gin engine, register core handlers,
		// actuator endpoints and application routers
		fx.Provide(gin.New),
		fx.Invoke(router.RegisterGinRouters),
	}
}
