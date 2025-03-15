package events

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Publisher interface {
	Publish(ctx context.Context, e Event) error
}

type RedisPublisher struct {
	client *redis.Client
}

type Payload struct {
	EventType string `json:"event_type"`
	Data      Event  `json:"data"`
}

func NewPublisher(redisAddr string) *RedisPublisher {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	return &RedisPublisher{
		client: client,
	}
}

func (p *RedisPublisher) Publish(ctx context.Context, e Event) error {
	payload := Payload{
		EventType: e.getEventType(),
		Data:      e,
	}
	err := p.client.Publish(ctx, "New Event: ", payload).Err()
	if err != nil {
		return err
	}
	return nil
}
