package queues

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/quanndh/go-app/public/config"
	"go.uber.org/fx"
)

type QueueClientIn struct {
	fx.In
	Config *config.Configuration
}

func NewQueueClient(q QueueClientIn) *asynq.Client {
	redisAddr := fmt.Sprintf("%s:%d", q.Config.App.Redis.Host, q.Config.App.Redis.Port)
	fmt.Println(redisAddr)
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})

	return client
}
