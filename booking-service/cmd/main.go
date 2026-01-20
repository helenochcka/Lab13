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

	accountCli, _ := client.NewAccountClient("proto-service:50051")
	eventCli, _ := client.NewEventClient("proto-service:50052")
	notifyCli, _ := client.NewNotificationClient("proto-service:50054")

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
