package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Account struct {
	ID    string
	Email string
	Name  string
}

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(ctx context.Context, email, name string) (*Account, error) {
	id := uuid.New().String()

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO accounts (id, email, name) VALUES ($1, $2, $3)`,
		id, email, name,
	)
	if err != nil {
		return nil, err
	}

	return &Account{ID: id, Email: email, Name: name}, nil
}

func (r *AccountRepository) Get(ctx context.Context, id string) (*Account, error) {
	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, email, name FROM accounts WHERE id=$1`, id,
	)

	var acc Account
	if err := row.Scan(&acc.ID, &acc.Email, &acc.Name); err != nil {
		return nil, err
	}

	return &acc, nil
}

func (r *AccountRepository) Exists(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(
		ctx,
		`SELECT EXISTS(SELECT 1 FROM accounts WHERE id=$1)`, id,
	).Scan(&exists)

	return exists, err
}
