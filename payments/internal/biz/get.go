package biz
import (
	"context"
	pb "payments/api/payments/v1"
	"payments/helper"

	"payments/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func (uc *PaymentUsecase) Get(
	ctx context.Context,
	req *pb.GetPaymentRequest,
) (*pb.PaymentResponse, error) {

	// helper validation
	if err := helper.ValidateGetPayment(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// data call
	payment, err := uc.repo.GetDB(ctx, req.PaymentId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "payment not found")
	}

	// mapper (model-response)
	return mapper.BizToPaymentResponse(payment), nil
}
