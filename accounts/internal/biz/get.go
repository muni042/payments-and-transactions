package biz
import (
	"context"

	pb "accounts/api/accounts/v1"
	"accounts/internal/helper"
	"accounts/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GET ACCOUNT
func (as *AccountUsecase) Get(
	ctx context.Context,
	req *pb.GetAccountRequest,
) (*pb.AccountResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	account, err := as.repo.GetDB(ctx, req.AccountId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	return mapper.BizToAccountResponse(account), nil
}
