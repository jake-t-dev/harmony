package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type Message interface {
	CreateMessage(message *Message) error
	GetMessage(messageUUID uuid.UUID) (*Message, error)
	GetMessages() ([]*Message, error)
	UpdateMessage(message *Message) error
	DeleteMessage(messageUUID uuid.UUID) error
}

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) CreateMessage(message *Message) error {
	return nil
}

func (r *MessageRepo) GetMessage(messageUUID uuid.UUID) (*Message, error) {
	return nil, nil
}

func (r *MessageRepo) GetMessages() ([]*Message, error) {
	return nil, nil
}

func (r *MessageRepo) UpdateMessage(message *Message) error {
	return nil
}

func (r *MessageRepo) DeleteMessage(messageUUID uuid.UUID) error {
	return nil
}
