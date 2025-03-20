package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type DirectMessage interface {
	SendDirectMessage(directMessage *DirectMessage) error
	GetDirectMessagesBySenderID(senderID uuid.UUID) ([]*DirectMessage, error)
	GetDirectMessagesByReceiverID(receiverID uuid.UUID) ([]*DirectMessage, error)
	GetDirectMessage(directMessageID uuid.UUID) (*DirectMessage, error)
	UpdateDirectMessage(directMessage *DirectMessage) error
	DeleteDirectMessage(directMessageID uuid.UUID) error
}

type DirectMessageRepo struct {
	db *sql.DB
}

func NewDirectMessageRepo(db *sql.DB) *DirectMessageRepo {
	return &DirectMessageRepo{db: db}
}

func (r *DirectMessageRepo) SendDirectMessage(directMessage *DirectMessage) error {
	return nil
}

func (r *DirectMessageRepo) GetDirectMessage(directMessageID uuid.UUID) (*DirectMessage, error) {
	return nil, nil
}

func (r *DirectMessageRepo) GetDirectMessagesBySenderID(senderID uuid.UUID) ([]*DirectMessage, error) {
	return nil, nil
}

func (r *DirectMessageRepo) GetDirectMessagesByReceiverID(receiverID uuid.UUID) ([]*DirectMessage, error) {
	return nil, nil
}

func (r *DirectMessageRepo) UpdateDirectMessage(directMessage *DirectMessage) error {
	return nil
}

func (r *DirectMessageRepo) DeleteDirectMessage(directMessageID uuid.UUID) error {
	return nil
}
