package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type GroupMember interface {
	AddMember(groupID uuid.UUID, userID uuid.UUID) error
	RemoveMember(groupID uuid.UUID, userID uuid.UUID) error
}

type GroupMemberRepo struct {
	db *sql.DB
}

func NewGroupMemberRepo(db *sql.DB) *GroupMemberRepo {
	return &GroupMemberRepo{db: db}
}

func (r *GroupMemberRepo) AddMember(groupID uuid.UUID, userID uuid.UUID) error {
	return nil
}

func (r *GroupMemberRepo) RemoveMember(groupID uuid.UUID, userID uuid.UUID) error {
	return nil
}
