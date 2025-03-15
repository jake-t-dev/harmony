package events

import (
	"context"
	"encoding/json"
	"fmt"

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

func NewPublisher(redisAddr string) (*RedisPublisher, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Redis on ", redisAddr)

	return &RedisPublisher{
		client: client,
	}, nil

}

func (p *RedisPublisher) Close() error {
	return p.client.Close()
}

func (p *RedisPublisher) Publish(ctx context.Context, e Event) error {
	payload := Payload{
		EventType: e.GetEventType(),
		Data:      e,
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error serializing event:", err)
		return err
	}

	err = p.client.Publish(ctx, "events", payloadData).Err()
	if err != nil {
		return err
	}
	return nil
}
