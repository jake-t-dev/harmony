package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type Channel interface {
	CreateChannel(channel *Channel) error
	GetChannel(channelID uuid.UUID) (*Channel, error)
	GetChannelsByGroupID(groupID uuid.UUID) ([]*Channel, error)
	UpdateChannel(channel *Channel) error
	DeleteChannel(channelID uuid.UUID) error
}

type ChannelRepo struct {
	db *sql.DB
}

func NewChannelRepo(db *sql.DB) *ChannelRepo {
	return &ChannelRepo{db: db}
}

func (r *ChannelRepo) CreateChannel(channel *Channel) error {
	return nil
}

func (r *ChannelRepo) GetChannel(channelID uuid.UUID) (*Channel, error) {
	return nil, nil
}

func (r *ChannelRepo) GetChannelsByGroupID(groupID uuid.UUID) ([]*Channel, error) {
	return nil, nil
}

func (r *ChannelRepo) UpdateChannel(channel *Channel) error {
	return nil
}

func (r *ChannelRepo) DeleteChannel(channelID uuid.UUID) error {
	return nil
}
