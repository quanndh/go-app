package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/quanndh/go-app/adapter/queues"
)

type QueueHandler struct {
}

func NewQueueHandler() *QueueHandler {
	return &QueueHandler{}
}

func (h QueueHandler) HandleCreatedUser(ctx context.Context, t *asynq.Task) error {
	var p queues.CreatedUserPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	fmt.Printf("New user created: user_id=%d\n", p.UserId)
	// Email delivery code ...
	return nil
}
