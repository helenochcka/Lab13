package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type Event struct {
	ID             string
	Title          string
	TotalSeats     int32
	AvailableSeats int32
}

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Create(
	ctx context.Context,
	title string,
	totalSeats int32,
) (*Event, error) {

	id := uuid.New().String()

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO events (id, title, total_seats, available_seats)
		 VALUES ($1, $2, $3, $3)`,
		id, title, totalSeats,
	)
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:             id,
		Title:          title,
		TotalSeats:     totalSeats,
		AvailableSeats: totalSeats,
	}, nil
}

func (r *EventRepository) Get(ctx context.Context, id string) (*Event, error) {
	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, title, total_seats, available_seats FROM events WHERE id=$1`,
		id,
	)

	var e Event
	if err := row.Scan(&e.ID, &e.Title, &e.TotalSeats, &e.AvailableSeats); err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *EventRepository) List(ctx context.Context) ([]*Event, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, title, total_seats, available_seats FROM events`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.TotalSeats, &e.AvailableSeats); err != nil {
			return nil, err
		}
		events = append(events, &e)
	}

	return events, nil
}

func (r *EventRepository) ReserveSeat(ctx context.Context, eventID string) error {
	res, err := r.db.ExecContext(
		ctx,
		`UPDATE events
		 SET available_seats = available_seats - 1
		 WHERE id = $1 AND available_seats > 0`,
		eventID,
	)
	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("no available seats")
	}

	return nil
}

func (r *EventRepository) ReleaseSeat(ctx context.Context, eventID string) error {
	_, err := r.db.ExecContext(
		ctx,
		`UPDATE events
		 SET available_seats = available_seats + 1
		 WHERE id = $1 AND available_seats < total_seats`,
		eventID,
	)
	return err
}
