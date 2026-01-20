package client

import (
	"context"
	"time"

	accountpb "booking-service/proto/account"

	"google.golang.org/grpc"
)

type AccountClient struct {
	client accountpb.AccountServiceClient
}

func NewAccountClient(addr string) (*AccountClient, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(3*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return &AccountClient{
		client: accountpb.NewAccountServiceClient(conn),
	}, nil
}

func (c *AccountClient) Exists(ctx context.Context, accountID string) (bool, error) {
	resp, err := c.client.CheckAccountExists(ctx, &accountpb.CheckAccountExistsRequest{
		Id: accountID,
	})
	if err != nil {
		return false, err
	}

	return resp.Exists, nil
}
