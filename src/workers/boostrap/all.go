package boostrap

import (
	"github.com/quanndh/go-app/worker/config"
	"github.com/quanndh/go-app/worker/handlers"
	"github.com/quanndh/go-app/worker/queues"
	"go.uber.org/fx"
)

func All() []fx.Option {
	return []fx.Option{
		fx.Provide(config.LoadConfig),

		fx.Provide(queues.NewQueueServer),

		fx.Provide(handlers.NewQueueHandler),

		fx.Invoke(queues.RegisterHandlers),
	}
}
