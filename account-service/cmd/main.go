package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"account-service/internal/db"
	"account-service/internal/repository"
	grpcHandler "account-service/internal/transport/grpc"

	accountpb "account-service/proto"
)

func main() {
	dsn := os.Getenv("POSTGRES_DSN")

	dbConn, err := db.NewPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewAccountRepository(dbConn)
	handler := grpcHandler.NewAccountHandler(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	accountpb.RegisterAccountServiceServer(server, handler)

	log.Println("Account Service started on :50051")
	log.Fatal(server.Serve(lis))
}
