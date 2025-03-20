package models

import "harmony-storage/events"

// GroupInviteEvent is a struct that represents a group invite event
type GroupInviteEvent struct {
	GroupInvite events.GroupInvite `json:"group_invite"`
}

func (e GroupInviteEvent) GetEventType() string {
	return "group_invite"
}

// GroupInviteAcceptEvent is a struct that represents a group invite accept event
type GroupInviteAcceptEvent struct {
	GroupInvite events.GroupInvite `json:"group_invite"`
}

func (e GroupInviteAcceptEvent) GetEventType() string {
	return "group_invite_accept"
}

// GroupInviteDeleteEvent is a struct that represents a group invite delete event
type GroupInviteDeleteEvent struct {
	GroupInvite events.GroupInvite `json:"group_invite"`
}
