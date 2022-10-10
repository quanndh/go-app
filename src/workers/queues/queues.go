package queues

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/quanndh/go-app/adapter/queues"
	"github.com/quanndh/go-app/worker/config"
	"github.com/quanndh/go-app/worker/handlers"
	"go.uber.org/fx"
	"log"
)

type QueueServerIn struct {
	fx.In
	Config *config.Configuration
}

func NewQueueServer(q QueueServerIn) *asynq.Server {
	redisAddr := fmt.Sprintf("%s:%d", q.Config.App.Redis.Host, q.Config.App.Redis.Port)
	fmt.Println(redisAddr)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	return srv
}

type RegisterHandlerIn struct {
	fx.In

	QueueServer *asynq.Server
	Handler     *handlers.QueueHandler
}

func RegisterHandlers(p RegisterHandlerIn) {
	mux := asynq.NewServeMux()

	mux.HandleFunc(queues.CreatedUserQueue, p.Handler.HandleCreatedUser)

	if err := p.QueueServer.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
