package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	grpcHandler "notification-service/internal/transport/grpc"

	notificationpb "notification-service/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	handler := grpcHandler.NewNotificationHandler()

	notificationpb.RegisterNotificationServiceServer(server, handler)

	log.Println("Notification Service started on :50054")
	log.Fatal(server.Serve(lis))
}
