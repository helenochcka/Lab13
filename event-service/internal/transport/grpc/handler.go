package grpc

import (
	"context"

	"event-service/internal/repository"

	eventpb "event-service/proto"
)

type EventHandler struct {
	eventpb.UnimplementedEventServiceServer
	repo *repository.EventRepository
}

func NewEventHandler(repo *repository.EventRepository) *EventHandler {
	return &EventHandler{repo: repo}
}

func (h *EventHandler) CreateEvent(
	ctx context.Context,
	req *eventpb.CreateEventRequest,
) (*eventpb.CreateEventResponse, error) {

	event, err := h.repo.Create(ctx, req.Title, req.TotalSeats)
	if err != nil {
		return nil, err
	}

	return &eventpb.CreateEventResponse{
		Event: toProto(event),
	}, nil
}

func (h *EventHandler) GetEvent(
	ctx context.Context,
	req *eventpb.GetEventRequest,
) (*eventpb.GetEventResponse, error) {

	event, err := h.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &eventpb.GetEventResponse{
		Event: toProto(event),
	}, nil
}

func (h *EventHandler) ListEvents(
	ctx context.Context,
	_ *eventpb.ListEventsRequest,
) (*eventpb.ListEventsResponse, error) {

	events, err := h.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var protoEvents []*eventpb.Event
	for _, e := range events {
		protoEvents = append(protoEvents, toProto(e))
	}

	return &eventpb.ListEventsResponse{
		Events: protoEvents,
	}, nil
}

func (h *EventHandler) ReserveSeat(
	ctx context.Context,
	req *eventpb.ReserveSeatRequest,
) (*eventpb.ReserveSeatResponse, error) {

	err := h.repo.ReserveSeat(ctx, req.EventId)
	if err != nil {
		return &eventpb.ReserveSeatResponse{Success: false}, nil
	}

	return &eventpb.ReserveSeatResponse{Success: true}, nil
}

func (h *EventHandler) ReleaseSeat(
	ctx context.Context,
	req *eventpb.ReleaseSeatRequest,
) (*eventpb.ReleaseSeatResponse, error) {

	err := h.repo.ReleaseSeat(ctx, req.EventId)
	if err != nil {
		return &eventpb.ReleaseSeatResponse{Success: false}, nil
	}

	return &eventpb.ReleaseSeatResponse{Success: true}, nil
}

func toProto(e *repository.Event) *eventpb.Event {
	return &eventpb.Event{
		Id:             e.ID,
		Title:          e.Title,
		TotalSeats:     e.TotalSeats,
		AvailableSeats: e.AvailableSeats,
	}
}
