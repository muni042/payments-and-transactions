package biz

import (
	pb "accounts/api/accounts/v1"
	"accounts/internal/helper"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DEBIT
func (as *AccountUsecase) Debit(
	ctx context.Context,
	req *pb.DebitAccountRequest,
) (*pb.OperationResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.Amount <= 0 {
		return nil, status.Error(codes.InvalidArgument, "amount must be greater than zero")
	}

	ok, err := as.repo.Debit(ctx, req.AccountId, float64(req.Amount))
	if err != nil {
		return nil, status.Error(codes.Internal, "debit failed")
	}

	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "insufficient balance")
	}

	return &pb.OperationResponse{Success: true}, nil
}
