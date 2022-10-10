package queues

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	CreatedUserQueue = "user:created"
)

type CreatedUserPayload struct {
	UserId uint
}

func NewCreatedUserTask(userId uint) (*asynq.Task, error) {
	payload, err := json.Marshal(CreatedUserPayload{UserId: userId})

	if err != nil {
		return nil, err
	}
	return asynq.NewTask(CreatedUserQueue, payload), nil
}
