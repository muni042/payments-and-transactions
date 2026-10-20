package service

import (
	"context"

	pb "accounts/api/accounts/v1"
	"accounts/internal/biz"
)

type AccountsServiceService struct {
	pb.UnimplementedAccountsServiceServer
	uc *biz.AccountUsecase
}

func NewAccountsServiceService(uc *biz.AccountUsecase) *AccountsServiceService {
	return &AccountsServiceService{uc: uc}
}

func (s *AccountsServiceService) CreateAccount(
	ctx context.Context,
	req *pb.CreateAccountRequest,
) (*pb.AccountResponse, error) {
	return s.uc.Create(ctx, req)
}

func (s *AccountsServiceService) GetAccount(
	ctx context.Context,
	req *pb.GetAccountRequest,
) (*pb.AccountResponse, error) {
	return s.uc.Get(ctx, req)
}

func (s *AccountsServiceService) UpdateAccount(
	ctx context.Context,
	req *pb.UpdateAccountRequest,
) (*pb.AccountResponse, error) {
	return s.uc.Update(ctx, req)
}

func (s *AccountsServiceService) DeleteAccount(
	ctx context.Context,
	req *pb.DeleteAccountRequest,
) (*pb.DeleteAccountResponse, error) {
	return s.uc.Delete(ctx, req)
}

/*func (s *AccountsServiceService) ListAccount(
    ctx context.Context,
    req *pb.ListAccountRequest,
) (*pb.ListAccountResponse, error) {
    return s.uc.List(ctx, req)
}*/

// ✅ ACCOUNT EXISTS IMPLEMENTATION
func (s *AccountsServiceService) CheckAccountExists(
	ctx context.Context,
	req *pb.GetAccountRequest,
) (*pb.CheckAccountExistsResponse, error) {
	return s.uc.Exists(ctx, req)
}
func (s *AccountsServiceService) CreditAccount(
	ctx context.Context,
	req *pb.CreditAccountRequest,
) (*pb.OperationResponse, error) {
	return s.uc.Credit(ctx, req)
}

func (s *AccountsServiceService) DebitAccount(
	ctx context.Context,
	req *pb.DebitAccountRequest,
) (*pb.OperationResponse, error) {
	return s.uc.Debit(ctx, req)
}
