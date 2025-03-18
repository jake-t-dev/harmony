package repo

import (
	"database/sql"

	"github.com/google/uuid"
)

type User interface {
	CreateUser(user *User) error
	GetUser(userUUID uuid.UUID) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(userUUID uuid.UUID) error
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *User) error {
	return nil
}

func (r *UserRepo) GetUser(userUUID uuid.UUID) (*User, error) {
	return nil, nil
}

func (r *UserRepo) GetUsers() ([]*User, error) {
	return nil, nil
}

func (r *UserRepo) UpdateUser(user *User) error {
	return nil
}

func (r *UserRepo) DeleteUser(userUUID uuid.UUID) error {
	return nil
}
