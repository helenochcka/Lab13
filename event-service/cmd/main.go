package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"event-service/internal/db"
	"event-service/internal/repository"
	grpcHandler "event-service/internal/transport/grpc"
	eventpb "event-service/proto"
)

func main() {
	dsn := os.Getenv("POSTGRES_DSN")

	dbConn, err := db.NewPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewEventRepository(dbConn)
	handler := grpcHandler.NewEventHandler(repo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	eventpb.RegisterEventServiceServer(server, handler)

	log.Println("Event Service started on :50052")
	log.Fatal(server.Serve(lis))
}
