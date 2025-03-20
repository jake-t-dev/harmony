package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type GroupMessage interface {
	SendGroupMessage(groupMessage *GroupMessage) error
	GetGroupMessagesByGroupID(channelID uuid.UUID) ([]*GroupMessage, error)
	GetGroupMessagesByChannelID(channelID uuid.UUID) ([]*GroupMessage, error)
	GetGroupMessagesByUserID(userID uuid.UUID) ([]*GroupMessage, error)
	GetGroupMessage(groupMessageID uuid.UUID) (*GroupMessage, error)
	UpdateGroupMessage(groupMessage *GroupMessage) error
	DeleteGroupMessage(groupMessageID uuid.UUID) error
}

type GroupMessageRepo struct {
	db *sql.DB
}

func NewGroupMessageRepo(db *sql.DB) *GroupMessageRepo {
	return &GroupMessageRepo{db: db}
}

func (r *GroupMessageRepo) SendGroupMessage(groupMessage *GroupMessage) error {
	return nil
}

func (r *GroupMessageRepo) GetGroupMessage(groupMessageID uuid.UUID) (*GroupMessage, error) {
	return nil, nil
}

func (r *GroupMessageRepo) GetGroupMessagesByChannelID(channelID uuid.UUID) ([]*GroupMessage, error) {
	return nil, nil
}

func (r *GroupMessageRepo) GetGroupMessagesByUserID(userID uuid.UUID) ([]*GroupMessage, error) {
	return nil, nil
}

func (r *GroupMessageRepo) UpdateGroupMessage(groupMessage *GroupMessage) error {
	return nil
}

func (r *GroupMessageRepo) DeleteGroupMessage(groupMessageID uuid.UUID) error {
	return nil
}
