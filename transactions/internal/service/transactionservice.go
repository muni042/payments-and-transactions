package service

import (
	"context"

	pb "transactions/api/transactions/v1"
	"transactions/internal/biz"
)

type TransactionServiceService struct {
	pb.UnimplementedTransactionServiceServer
	uc *biz.TransactionUsecase
}

// Constructor
func NewTransactionServiceService(
	uc *biz.TransactionUsecase,
) *TransactionServiceService {
	return &TransactionServiceService{
		uc: uc,
	}
}

func (s *TransactionServiceService) AddTransaction(
	ctx context.Context,
	req *pb.AddTransactionRequest,
) (*pb.AddTransactionResponse, error) {

	return s.uc.Add(ctx, req)
}

func (s *TransactionServiceService) GetTransaction(
	ctx context.Context,
	req *pb.GetTransactionRequest,
) (*pb.GetTransactionResponse, error) {

	return s.uc.Get(ctx, req)
}

func (s *TransactionServiceService) UpdateTransaction(
	ctx context.Context,
	req *pb.UpdateTransactionRequest,
) (*pb.UpdateTransactionResponse, error) {

	return s.uc.Update(ctx, req)
}
