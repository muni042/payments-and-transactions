package biz

import (
	"context"

	pb "transactions/api/transactions/v1"

	"transactions/internal/helper"
	"transactions/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ADD TRANSACTION
func (tu *TransactionUsecase) Add(
	ctx context.Context,
	req *pb.AddTransactionRequest,
) (*pb.AddTransactionResponse, error) {

	// mapper (request → model)
	tx, err := mapper.AddReqToBiz(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// validation (helper)
	if err := helper.ValidateTransaction(tx); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// data call
	if err := tu.repo.Create(ctx, tx); err != nil {
		return nil, status.Error(codes.Internal, "failed to create transaction")
	}

	//mapper (model → response)
	return mapper.ModelToAddResponse(tx), nil
}
