package biz
import (
	"context"

	pb "accounts/api/accounts/v1"
	"accounts/internal/helper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
// CHECK ACCOUNT EXISTS
func (as *AccountUsecase) Exists(
	ctx context.Context,
	req *pb.GetAccountRequest,
) (*pb.CheckAccountExistsResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ok, err := as.repo.ExistsByID(ctx, req.AccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to check account")
	}

	return &pb.CheckAccountExistsResponse{Exists: ok}, nil
}