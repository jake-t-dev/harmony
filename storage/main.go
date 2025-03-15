package main

import (
	"context"
	"fmt"
	"harmony-storage/events"
	"harmony-storage/events/models"
	"harmony-storage/pkg/db"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	store, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	defer store.Close()

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	publisher, err := events.NewPublisher("localhost:6379")
	if err != nil {
		log.Fatal(err)
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
