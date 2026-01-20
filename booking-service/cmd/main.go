package main

import (
	bookingpb "booking-service/proto/booking"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"booking-service/internal/client"
	"booking-service/internal/db"
	"booking-service/internal/repository"
	grpcHandler "booking-service/internal/transport/grpc"
)

func main() {
	dbConn, err := db.NewPostgres(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	accountCli, err := client.NewAccountClient("account-service:50051")
	if err != nil {
		log.Fatalf("failed to create account client: %v", err)
	}
	eventCli, err := client.NewEventClient("event-service:50052")
	if err != nil {
		log.Fatalf("failed to create event client: %v", err)
	}
	notifyCli, err := client.NewNotificationClient("notification-service:50054")
	if err != nil {
		log.Fatalf("failed to create notification client: %v", err)
	}

	repo := repository.NewBookingRepository(dbConn)
	handler := grpcHandler.NewBookingHandler(repo, accountCli, eventCli, notifyCli)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	bookingpb.RegisterBookingServiceServer(server, handler)

	log.Println("Booking Service started on :50053")
	log.Fatal(server.Serve(lis))
}
