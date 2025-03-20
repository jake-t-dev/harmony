package db

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgresDB struct {
	db *sql.DB
}
type Storage interface {
	Init() error
	Close() error
	CreateTable(q string) error
}

func NewPostgresDB() (*PostgresDB, error) {
	connStr := "user=postgres dbname=postgres host=127.0.0.1 port=5432 password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to Postgres on localhost:5432")

	return &PostgresDB{db: db}, nil
}

func (s *PostgresDB) Init() error {
	queries := GetCreateDatabase()
	for _, q := range queries.Queries {
		if err := s.CreateTable(q.Query); err != nil {
			return err
		}
	}

	return nil
}

func (s *PostgresDB) Close() error {
	return s.db.Close()
}

func (s *PostgresDB) CreateTable(q string) error {
	_, err := s.db.Exec(q)
	return err
}
