package biz
import (
	"context"

	pb "accounts/api/accounts/v1"
	"accounts/internal/helper"
	"accounts/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
// UPDATE ACCOUNT
func (as *AccountUsecase) Update(
	ctx context.Context,
	req *pb.UpdateAccountRequest,
) (*pb.AccountResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	account := mapper.UpdateReqtoBiz(req)

	if err := as.repo.UpdateDB(ctx, account); err != nil {
		return nil, status.Error(codes.Internal, "account update failed")
	}

	account, _ = as.repo.GetDB(ctx, req.AccountId)
	return mapper.BizToAccountResponse(account), nil
}