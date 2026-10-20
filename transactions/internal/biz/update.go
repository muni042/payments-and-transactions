package biz
import (
	"context"
	

	pb "transactions/api/transactions/v1"
	
	"transactions/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
// UPDATE TRANSACTION
func (tu *TransactionUsecase) Update(
	ctx context.Context,
	req *pb.UpdateTransactionRequest,
) (*pb.UpdateTransactionResponse, error) {

	//mapper
	tx, err := mapper.UpdateReqToBiz(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//business validation
	if tx.Status == "" && tx.Description == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"either status or description must be provided",
		)
	}

	// data call
	if err := tu.repo.Update(ctx, tx); err != nil {
		return nil, status.Error(codes.Internal, "failed to update transaction")
	}

	// mapper
	return mapper.ModelToUpdateResponse(tx), nil
}