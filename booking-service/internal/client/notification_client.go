package client

import (
	"context"
	"time"

	notificationpb "github.com/helenochcka/project-proto/account"

	"google.golang.org/grpc"
)

type NotificationClient struct {
	client notificationpb.NotificationServiceClient
}

func NewNotificationClient(addr string) (*NotificationClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return &NotificationClient{
		client: notificationpb.NewNotificationServiceClient(conn),
	}, nil
}

func (c *NotificationClient) SendBookingCreated(ctx context.Context, bookingID string) error {
	_, err := c.client.SendBookingConfirmation(ctx, &notificationpb.BookingNotificationRequest{
		BookingId: bookingID,
	})
	return err
}
