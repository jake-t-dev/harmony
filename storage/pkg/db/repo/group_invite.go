package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type GroupInvite interface {
	CreateGroupInvite(groupInvite *GroupInvite) error
	GetGroupInvite(groupInviteID uuid.UUID) (*GroupInvite, error)
	GetGroupInvitesByGroupID(groupID uuid.UUID) ([]*GroupInvite, error)
	GetGroupInvitesBySenderID(senderID uuid.UUID) ([]*GroupInvite, error)
	GetGroupInvitesByReceiverID(receiverID uuid.UUID) ([]*GroupInvite, error)
	DeleteGroupInvite(groupInviteID uuid.UUID) error
}

type GroupInviteRepo struct {
	db *sql.DB
}

func NewGroupInviteRepo(db *sql.DB) *GroupInviteRepo {
	return &GroupInviteRepo{db: db}
}

func (r *GroupInviteRepo) CreateGroupInvite(groupInvite *GroupInvite) error {
	return nil
}

func (r *GroupInviteRepo) GetGroupInvite(groupInviteID uuid.UUID) (*GroupInvite, error) {
	return nil, nil
}

func (r *GroupInviteRepo) GetGroupInvitesByGroupID(groupID uuid.UUID) ([]*GroupInvite, error) {
	return nil, nil
}

func (r *GroupInviteRepo) GetGroupInvitesBySenderID(senderID uuid.UUID) ([]*GroupInvite, error) {
	return nil, nil
}

func (r *GroupInviteRepo) GetGroupInvitesByReceiverID(receiverID uuid.UUID) ([]*GroupInvite, error) {
	return nil, nil
}

func (r *GroupInviteRepo) DeleteGroupInvite(groupInviteID uuid.UUID) error {
	return nil
}
