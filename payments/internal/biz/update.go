package biz

import (
	"context"
	pb "payments/api/payments/v1"
	"payments/helper"

	"payments/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc *PaymentUsecase) Update(
	ctx context.Context,
	req *pb.UpdatePaymentRequest,
) (*pb.PaymentResponse, error) {

	// helper validation
	if err := helper.ValidateUpdatePayment(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	payment := mapper.UpdatePaymentReqToBiz(req)

	// data call
	if err := uc.repo.UpdateDB(ctx, payment); err != nil {
		return nil, status.Error(codes.Internal, "payment update failed")
	}

	// fetch updated record
	updated, err := uc.repo.GetDB(ctx, payment.PaymentID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "payment not found after update")
	}

	return mapper.BizToPaymentResponse(updated), nil
}
