package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type Group interface {
	CreateGroup(group *Group) error
	GetGroup(groupUUID uuid.UUID) (*Group, error)
	GetGroupsByUserID(userID uuid.UUID) ([]*Group, error)
	GetGroups() ([]*Group, error)
	UpdateGroup(group *Group) error
	DeleteGroup(groupUUID uuid.UUID) error
}

type GroupRepo struct {
	db *sql.DB
}

func NewGroupRepo(db *sql.DB) *GroupRepo {
	return &GroupRepo{db: db}
}

func (r *GroupRepo) CreateGroup(group *Group) error {
	return nil
}

func (r *GroupRepo) GetGroup(groupUUID uuid.UUID) (*Group, error) {
	return nil, nil
}

func (r *GroupRepo) GetGroupsByUserID(userID uuid.UUID) ([]*Group, error) {
	return nil, nil
}

func (r *GroupRepo) GetGroups() ([]*Group, error) {
	return nil, nil
}

func (r *GroupRepo) UpdateGroup(group *Group) error {
	return nil
}

func (r *GroupRepo) DeleteGroup(groupUUID uuid.UUID) error {
	return nil
}
