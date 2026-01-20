package grpc

import (
	"context"
	"log"
	"time"

	notificationpb "notification-service/proto"
)

type NotificationHandler struct {
	notificationpb.UnimplementedNotificationServiceServer
}

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{}
}

func (h *NotificationHandler) SendBookingConfirmation(
	ctx context.Context,
	req *notificationpb.BookingNotificationRequest,
) (*notificationpb.BookingNotificationResponse, error) {

	time.Sleep(300 * time.Millisecond)

	log.Printf(
		"[Notification] Booking confirmation sent for booking_id=%s",
		req.BookingId,
	)

	return &notificationpb.BookingNotificationResponse{
		Success: true,
	}, nil
}

func (h *NotificationHandler) SendBookingCancellation(
	ctx context.Context,
	req *notificationpb.BookingNotificationRequest,
) (*notificationpb.BookingNotificationResponse, error) {

	log.Printf(
		"[Notification] Booking cancellation sent for booking_id=%s",
		req.BookingId,
	)

	return &notificationpb.BookingNotificationResponse{
		Success: true,
	}, nil
}
