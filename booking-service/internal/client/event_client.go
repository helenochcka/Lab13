package client

import (
	"context"
	"time"

	eventpb "booking-service/proto/event"

	"google.golang.org/grpc"
)

type EventClient struct {
	client eventpb.EventServiceClient
}

func NewEventClient(addr string) (*EventClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(3*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return &EventClient{
		client: eventpb.NewEventServiceClient(conn),
	}, nil
}

func (c *EventClient) ReserveSeat(ctx context.Context, eventID string) (bool, error) {
	resp, err := c.client.ReserveSeat(ctx, &eventpb.ReserveSeatRequest{
		EventId: eventID,
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (c *EventClient) ReleaseSeat(ctx context.Context, eventID string) {
	_, _ = c.client.ReleaseSeat(ctx, &eventpb.ReleaseSeatRequest{
		EventId: eventID,
	})
}
