package models

import "harmony-storage/events"

// FriendRequestEvent is a struct that represents a friend request event
type FriendRequestEvent struct {
	FriendRequest events.FriendRequest `json:"friend_request"`
}

func (e FriendRequestEvent) getEventType() string {
	return "friend_request"
}

// FriendRequestAcceptEvent is a struct that represents a friend request accept event
type FriendRequestAcceptEvent struct {
	FriendRequest events.FriendRequest `json:"friend_request"`
}

func (e FriendRequestAcceptEvent) getEventType() string {
	return "friend_request_accept"
}

// FriendRequestDeleteEvent is a struct that represents a friend request delete event
type FriendRequestDeleteEvent struct {
	FriendRequest events.FriendRequest `json:"friend_request"`
}

func (e FriendRequestDeleteEvent) getEventType() string {
	return "friend_request_delete"
}
