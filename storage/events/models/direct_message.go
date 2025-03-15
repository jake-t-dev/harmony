package models

import "harmony-storage/events"

// DirectMessageEvent is a struct that represents a direct message event
type DirectMessageEvent struct {
	Message   events.Message `json:"message"`
	Recipient events.User    `json:"recipient"`
}

func (e DirectMessageEvent) GetEventType() string {
	return "direct_message"
}

// DirectMessageUpdateEvent is a struct that represents a direct message update event
type DirectMessageUpdateEvent struct {
	Message events.Message `json:"message"`
}

func (e DirectMessageUpdateEvent) GetEventType() string {
	return "direct_message_update"
}

// DirectMessageDeleteEvent is a struct that represents a direct message delete event
type DirectMessageDeleteEvent struct {
	Message events.Message `json:"message"`
}

func (e DirectMessageDeleteEvent) GetEventType() string {
	return "direct_message_delete"
}
