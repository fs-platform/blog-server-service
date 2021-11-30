package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	blog "blog/proto/blog"
)

type Blog struct{}

func (e *Blog) Handle(ctx context.Context, msg *blog.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *blog.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
