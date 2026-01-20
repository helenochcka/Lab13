package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"booking-service/internal/client"
	"booking-service/internal/db"
	"booking-service/internal/repository"
	grpcHandler "booking-service/internal/transport/grpc"
	bookingpb "booking-service/proto"
)

func main() {
	dbConn, err := db.NewPostgres(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	accountCli, _ := client.NewAccountClient("account-service:50051")
	eventCli, _ := client.NewEventClient("event-service:50052")
	notifyCli, _ := client.NewNotificationClient("notification-service:50054")

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
