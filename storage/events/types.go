package events

import "github.com/google/uuid"

// Event is an interface that all events must implement
type Event interface {
	getEventType() string
}

// Group is a struct that represents a group chat
type Group struct {
	GroupID uuid.UUID
}

// Channel is a struct that represents a channel in a group chat
type Channel struct {
	Group     Group
	ChannelID uuid.UUID
}

// User is a struct that represents a user
type User struct {
	UserID uuid.UUID
}

// Message is a struct that represents a message containing a message ID and a sender
type Message struct {
	MessageID uuid.UUID
	Sender    User
}

// FriendRequest is a struct that represents a friend request containing a request ID, sender, and receiver
type FriendRequest struct {
	RequestID uuid.UUID
	Sender    User
	Receiver  User
}
