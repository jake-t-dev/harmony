package repo

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserQueries interface {
	CreateUser(user *User) error
	GetUser(userUUID uuid.UUID) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(userUUID uuid.UUID) error
}

type User struct {
	UUID         uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *User) error {
	query := `INSERT INTO users (id, username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, user.UUID, user.Name, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepo) GetUser(userUUID uuid.UUID) (*User, error) {
	query := `SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE id = $1`
	rows, err := r.db.Query(query, userUUID)
	if err != nil {
		return nil, err
	}

	user, err := scanIntoUser(rows)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetUsers() ([]*User, error) {
	query := `SELECT id, username, email, password_hash, created_at, updated_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (r *UserRepo) UpdateUser(user *User) error {
	query := `UPDATE users SET username = $1, email = $2, password_hash = $3, updated_at = $4 WHERE id = $5`
	_, err := r.db.Exec(query, user.Name, user.Email, user.PasswordHash, user.UpdatedAt, user.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteUser(userUUID uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, userUUID)
	if err != nil {
		return err
	}
	return nil
}

func scanIntoUser(rows *sql.Rows) (*User, error) {
	var user User
	err := rows.Scan(&user.UUID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
