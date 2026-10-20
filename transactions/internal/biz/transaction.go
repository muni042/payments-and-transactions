package biz

import (
	"transactions/internal/data"
)

type TransactionUsecase struct {
	repo *data.TransactionRepo
}

func NewTransactionUsecase(repo *data.TransactionRepo) *TransactionUsecase {
	return &TransactionUsecase{repo: repo}
}

/*
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

// GET TRANSACTION
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
*/
