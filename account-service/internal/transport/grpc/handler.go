package grpc

import (
	"context"

	"account-service/internal/repository"
	accountpb "account-service/proto"
)

type AccountHandler struct {
	accountpb.UnimplementedAccountServiceServer
	repo *repository.AccountRepository
}

func NewAccountHandler(repo *repository.AccountRepository) *AccountHandler {
	return &AccountHandler{repo: repo}
}

func (h *AccountHandler) CreateAccount(
	ctx context.Context,
	req *accountpb.CreateAccountRequest,
) (*accountpb.CreateAccountResponse, error) {

	acc, err := h.repo.Create(ctx, req.Email, req.Name)
	if err != nil {
		return nil, err
	}

	return &accountpb.CreateAccountResponse{
		Account: &accountpb.Account{
			Id:    acc.ID,
			Email: acc.Email,
			Name:  acc.Name,
		},
	}, nil
}

func (h *AccountHandler) GetAccount(
	ctx context.Context,
	req *accountpb.GetAccountRequest,
) (*accountpb.GetAccountResponse, error) {

	acc, err := h.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &accountpb.GetAccountResponse{
		Account: &accountpb.Account{
			Id:    acc.ID,
			Email: acc.Email,
			Name:  acc.Name,
		},
	}, nil
}

func (h *AccountHandler) CheckAccountExists(
	ctx context.Context,
	req *accountpb.CheckAccountExistsRequest,
) (*accountpb.CheckAccountExistsResponse, error) {

	exists, err := h.repo.Exists(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &accountpb.CheckAccountExistsResponse{
		Exists: exists,
	}, nil
}
