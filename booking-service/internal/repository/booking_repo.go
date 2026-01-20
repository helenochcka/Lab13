package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Booking struct {
	ID        string
	AccountID string
	EventID   string
	Status    string
}

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Create(
	ctx context.Context,
	accountID, eventID string,
) (*Booking, error) {

	id := uuid.New().String()

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO bookings (id, account_id, event_id, status)
		 VALUES ($1, $2, $3, $4)`,
		id, accountID, eventID, "CREATED",
	)
	if err != nil {
		return nil, err
	}

	return &Booking{
		ID:        id,
		AccountID: accountID,
		EventID:   eventID,
		Status:    "CREATED",
	}, nil
}

func (r *BookingRepository) Get(ctx context.Context, id string) (*Booking, error) {
	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, account_id, event_id, status FROM bookings WHERE id=$1`,
		id,
	)

	var b Booking
	if err := row.Scan(&b.ID, &b.AccountID, &b.EventID, &b.Status); err != nil {
		return nil, err
	}

	return &b, nil
}

func (r *BookingRepository) Cancel(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(
		ctx,
		`UPDATE bookings SET status='CANCELLED' WHERE id=$1`,
		id,
	)
	return err
}
