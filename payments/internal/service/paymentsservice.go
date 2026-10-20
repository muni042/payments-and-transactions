package service

import (
	"context"

	pb "payments/api/payments/v1"
	"payments/internal/biz"
)

type PaymentsServiceService struct {
	pb.UnimplementedPaymentsServiceServer
	uc *biz.PaymentUsecase
}

func NewPaymentsServiceService(uc *biz.PaymentUsecase) *PaymentsServiceService {
	return &PaymentsServiceService{uc: uc}
}

// create
func (s *PaymentsServiceService) CreatePayment(
	ctx context.Context,
	req *pb.CreatePaymentRequest,
) (*pb.PaymentResponse, error) {
	return s.uc.Create(ctx, req)
}

func (s *PaymentsServiceService) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{}, nil
}
func (s *PaymentsServiceService) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.PaymentResponse, error) {
	return &pb.PaymentResponse{}, nil
}
func (s *PaymentsServiceService) ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error) {
	return &pb.ListPaymentsResponse{}, nil
}

/*func (s *PaymentsServiceService) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
    return &pb.DeletePaymentResponse{}, nil
}
func (s *PaymentsServiceService) ListPayments(ctx context.Context, req *pb.ListPaymentsRequest) (*pb.ListPaymentsResponse, error) {
    return &pb.ListPaymentsResponse{}, nil
}
func (s *PaymentsServiceService) CheckPaymentExists(ctx context.Context, req *pb.GetPaymentRequest) (*pb.CheckPaymentExistsResponse, error) {
    return &pb.CheckPaymentExistsResponse{}, nil
}*/
