package biz
import (
	"context"
	"strconv"

	pb "transactions/api/transactions/v1"
	"transactions/internal/helper"
	"transactions/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func (tu *TransactionUsecase) Get(
	ctx context.Context,
	req *pb.GetTransactionRequest,
) (*pb.GetTransactionResponse, error) {

	// validation
	if err := helper.ValidateGetTransaction(req.TransactionId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//parse id
	transactionID, err := strconv.ParseInt(req.TransactionId, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid transaction_id")
	}

	//data call
	tx, err := tu.repo.GetByID(ctx, transactionID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	// mapper (model → response)
	return mapper.ModelToGetResponse(tx), nil
}
