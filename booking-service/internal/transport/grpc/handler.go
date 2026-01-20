package grpc

import (
	"context"
	"errors"
	"log"
	"time"

	bookingpb "booking-service/proto/booking"

	"booking-service/internal/client"
	"booking-service/internal/repository"
)

type BookingHandler struct {
	bookingpb.UnimplementedBookingServiceServer
	repo       *repository.BookingRepository
	accountCli *client.AccountClient
	eventCli   *client.EventClient
	notifyCli  *client.NotificationClient
}

func NewBookingHandler(
	repo *repository.BookingRepository,
	accountCli *client.AccountClient,
	eventCli *client.EventClient,
	notifyCli *client.NotificationClient,
) *BookingHandler {
	return &BookingHandler{
		repo:       repo,
		accountCli: accountCli,
		eventCli:   eventCli,
		notifyCli:  notifyCli,
	}
}

func (h *BookingHandler) CreateBooking(
	ctx context.Context,
	req *bookingpb.CreateBookingRequest,
) (*bookingpb.CreateBookingResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	exists, err := h.accountCli.Exists(ctx, req.AccountId)
	if err != nil || !exists {
		return nil, errors.New("proto not found")
	}

	ok, err := h.eventCli.ReserveSeat(ctx, req.EventId)
	if err != nil || !ok {
		return nil, errors.New("no seats available")
	}

	booking, err := h.repo.Create(ctx, req.AccountId, req.EventId)
	if err != nil {
		h.eventCli.ReleaseSeat(ctx, req.EventId)
		return nil, err
	}

	// best-effort proto
	go func() {
		if err := h.notifyCli.SendBookingCreated(context.Background(), booking.ID); err != nil {
			log.Println("proto failed:", err)
		}
	}()

	return &bookingpb.CreateBookingResponse{
		Booking: toProto(booking),
	}, nil
}

func (h *BookingHandler) GetBooking(
	ctx context.Context,
	req *bookingpb.GetBookingRequest,
) (*bookingpb.GetBookingResponse, error) {

	booking, err := h.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetBookingResponse{
		Booking: toProto(booking),
	}, nil
}

func (h *BookingHandler) CancelBooking(
	ctx context.Context,
	req *bookingpb.CancelBookingRequest,
) (*bookingpb.CancelBookingResponse, error) {

	err := h.repo.Cancel(ctx, req.Id)
	if err != nil {
		return &bookingpb.CancelBookingResponse{Success: false}, nil
	}

	return &bookingpb.CancelBookingResponse{Success: true}, nil
}

func toProto(b *repository.Booking) *bookingpb.Booking {
	return &bookingpb.Booking{
		Id:        b.ID,
		AccountId: b.AccountID,
		EventId:   b.EventID,
		Status:    b.Status,
	}
}
