package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // default DB
	})

	// Subscribe to the 'events' channel
	pubsub := rdb.Subscribe("events")
	defer pubsub.Close()

	// Wait for confirmation that subscription is active
	_, err := pubsub.Receive()
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	// Get the channel to receive messages
	ch := pubsub.Channel()

	fmt.Println("Listening for messages on channel 'events'...")

	for msg := range ch {
		fmt.Printf("Received message from channel [%s]: %s\n", msg.Channel, msg.Payload)
	}
}
