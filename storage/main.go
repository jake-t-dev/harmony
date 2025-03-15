package main

import (
	"context"
	"fmt"
	"harmony-storage/events"
	"harmony-storage/events/models"

	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	publisher, err := events.NewPublisher("localhost:6379")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer publisher.Close()

	newDirectMessage := models.DirectMessageEvent{
		Message: events.Message{
			MessageID: uuid.New(),
			Sender: events.User{
				UserID: uuid.New(),
			},
		},
		Recipient: events.User{
			UserID: uuid.New(),
		},
	}

	publisher.Publish(ctx, newDirectMessage)

	fmt.Println(ctx, publisher)
}
