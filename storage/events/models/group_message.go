package models

import "harmony-storage/events"

// GroupMessageEvent is a struct that represents a group message event
type GroupMessageEvent struct {
	Message events.Message `json:"message"`
	Channel events.Channel `json:"channel"`
}

func (e GroupMessageEvent) GetEventType() string {
	return "group_message"
}

// GroupMessageUpdateEvent is a struct that represents a group message update event
type GroupMessageUpdateEvent struct {
	Message events.Message `json:"message"`
}

func (e GroupMessageUpdateEvent) GetEventType() string {
	return "group_message_update"
}

// GroupMessageDeleteEvent is a struct that represents a group message delete event
type GroupMessageDeleteEvent struct {
	Message events.Message `json:"message"`
}

func (e GroupMessageDeleteEvent) GetEventType() string {
	return "group_message_delete"
}
