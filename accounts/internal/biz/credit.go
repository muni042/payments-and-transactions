package biz
import (
	"context"
	pb "accounts/api/accounts/v1"
	"accounts/internal/helper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
// CREDIT
func (as *AccountUsecase) Credit(
	ctx context.Context,
	req *pb.CreditAccountRequest,
) (*pb.OperationResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.Amount <= 0 {
		return nil, status.Error(codes.InvalidArgument, "amount must be greater than zero")
	}

	if err := as.repo.Credit(ctx, req.AccountId, float64(req.Amount)); err != nil {
		return nil, status.Error(codes.Internal, "credit failed")
	}

	return &pb.OperationResponse{Success: true}, nil
}