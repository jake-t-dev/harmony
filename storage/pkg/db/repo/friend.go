package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type Friend interface {
	CreateFriendRequest(userID uuid.UUID, friendID uuid.UUID) error
	UpdateFriendRequestStatus(userID uuid.UUID, friendID uuid.UUID, status string) error
	RemoveFriend(userID uuid.UUID, friendID uuid.UUID) error
}

type FriendRepo struct {
	db *sql.DB
}

func NewFriendRepo(db *sql.DB) *FriendRepo {
	return &FriendRepo{db: db}
}

func (r *FriendRepo) CreateFriendRequest(userID uuid.UUID, friendID uuid.UUID) error {
	return nil
}

func (r *FriendRepo) UpdateFriendRequestStatus(userID uuid.UUID, friendID uuid.UUID, status string) error {
	return nil
}

func (r *FriendRepo) RemoveFriend(userID uuid.UUID, friendID uuid.UUID) error {
	return nil
}
