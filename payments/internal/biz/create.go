package biz

import (
	accountspb "accounts/api/accounts/v1"
	"context"
	pb "payments/api/payments/v1"
	"payments/helper"

	"payments/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// create
func (uc *PaymentUsecase) Create(
	ctx context.Context,
	req *pb.CreatePaymentRequest,
) (*pb.PaymentResponse, error) {

	//Validate request
	if err := helper.ValidateCreatePayment(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// VERIFY CARD PIN
	if err := uc.verifyCardPin(ctx, req.FromAccountId, req.Pin); err != nil {
		return nil, err
	}

	//Map payment
	payment, err := mapper.CreatePaymentReqToBiz(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "payment mapping failed")
	}

	//Debit account
	debitResp, err := uc.account.DebitAccount(ctx, &accountspb.DebitAccountRequest{
		AccountId: req.FromAccountId,
		Amount:    int64(payment.Amount),
	})
	if err != nil {
		return nil, status.Error(codes.Unavailable, "account debit failed")
	}
	if !debitResp.Success {
		return nil, status.Error(codes.FailedPrecondition, "insufficient balance")
	}

	//Credit account
	creditResp, err := uc.account.CreditAccount(ctx, &accountspb.CreditAccountRequest{
		AccountId: req.ToAccountId,
		Amount:    int64(payment.Amount),
	})
	if err != nil {
		return nil, status.Error(codes.Unavailable, "account credit failed")
	}
	if !creditResp.Success {
		return nil, status.Error(codes.Internal, "credit operation failed")
	}

	//Save payment
	if err := uc.repo.InsertDB(ctx, payment); err != nil {
		return nil, status.Error(codes.Internal, "failed to save payment")
	}

	//Transaction DEBIT
	if _, err := uc.transaction.AddTransaction(
		ctx,
		mapper.PaymentToDebitTransaction(
			req.FromAccountId,
			req.Amount,
			payment.Currency,
		),
	); err != nil {
		return nil, status.Error(codes.Internal, "failed to record debit transaction")
	}

	// Transaction CREDIT
	if _, err := uc.transaction.AddTransaction(
		ctx,
		mapper.PaymentToCreditTransaction(
			req.ToAccountId,
			req.Amount,
			payment.Currency,
		),
	); err != nil {
		return nil, status.Error(codes.Internal, "failed to record credit transaction")
	}

	return mapper.BizToPaymentResponse(payment), nil
}
