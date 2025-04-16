package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//go:generate mockgen -destination=mocks/mock_user_cache.go -package=mocks crudTestTask/internal/repository DataBaseHandler
type DataBaseHandler interface {
	Create(user Data) (*Data, error)
	Get(id int64) (*Data, error)
	Update(user Data) (*Data, error)
	Delete(id int64) error
}

type Repository struct {
	db *sql.DB
}

func New(dbConn string) (*Repository, error) {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Successfully connected to database")

	return &Repository{db}, nil
}

func (r *Repository) Create(user Data) (*Data, error) {
	stmt := "INSERT INTO users (user_id, user_name) VALUES ($1, $2)"
	_, err := r.db.Exec(stmt, user.Id, user.Name)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return &user, nil
}

func (r *Repository) Get(id int64) (*Data, error) {
	stmt := "SELECT user_id, user_name FROM users WHERE user_id = $1"
	row := r.db.QueryRow(stmt, id)
	data := Data{}

	if err := row.Scan(&data.Id, &data.Name); err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return &data, nil
}

func (r *Repository) Update(user Data) (*Data, error) {
	stmt := "UPDATE users SET user_name = $1 WHERE user_id = $2"
	_, err := r.db.Exec(stmt, user.Name, user.Id)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return &user, nil
}

func (r *Repository) Delete(id int64) error {
	stmt := "DELETE FROM users WHERE user_id = $1"
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return fmt.Errorf("error delete user: %w", err)
	}

	return nil
}
