package biz
import (
	"context"

	pb "accounts/api/accounts/v1"
	
	"accounts/internal/helper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DELETE ACCOUNT
func (as *AccountUsecase) Delete(
	ctx context.Context,
	req *pb.DeleteAccountRequest,
) (*pb.DeleteAccountResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := as.repo.DeleteDB(ctx, req.AccountId); err != nil {
		return nil, status.Error(codes.Internal, "account deletion failed")
	}

	return &pb.DeleteAccountResponse{
		Message: "Account discarded successfully",
	}, nil
}
