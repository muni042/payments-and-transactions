package biz

import (
	accountspb "accounts/api/accounts/v1"
	cardspb "cards/api/cards/v1"
	"payments/internal/data"
	transactionspb "transactions/api/transactions/v1"
)

type PaymentUsecase struct {
	repo        *data.PaymentRepo
	account     accountspb.AccountsServiceClient
	transaction transactionspb.TransactionServiceClient
	cards       cardspb.CardsServiceClient
}

func NewPaymentUsecase(
	repo *data.PaymentRepo,
	account accountspb.AccountsServiceClient,
	transaction transactionspb.TransactionServiceClient,
	cards cardspb.CardsServiceClient,
) *PaymentUsecase {
	return &PaymentUsecase{
		repo:        repo,
		account:     account,
		transaction: transaction,
		cards:       cards,
	}
}

/*// create
func (uc *PaymentUsecase) Create(
	ctx context.Context,
	req *pb.CreatePaymentRequest,
) (*pb.PaymentResponse, error) {

	// 1. Validate request
	if err := helper.ValidateCreatePayment(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 2. VERIFY CARD PIN (BEFORE ANY MONEY MOVEMENT)
	if err := uc.verifyCardPin(ctx, req.FromAccountId, req.Pin); err != nil {
		return nil, err
	}

	// 3. Map payment
	payment, err := mapper.CreatePaymentReqToBiz(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "payment mapping failed")
	}

	// 4. Debit account
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

	// 5. Credit account
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

	// 6. Save payment
	if err := uc.repo.InsertDB(ctx, payment); err != nil {
		return nil, status.Error(codes.Internal, "failed to save payment")
	}

	// 7. Transaction (DEBIT)
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

	// 8. Transaction (CREDIT)
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

	// 9. Response
	return mapper.BizToPaymentResponse(payment), nil
}*/
/*
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

	// mapper (model → response)
	return mapper.BizToPaymentResponse(payment), nil
}

func (uc *PaymentUsecase) Update(
	ctx context.Context,
	req *pb.UpdatePaymentRequest,
) (*pb.PaymentResponse, error) {

	// helper validation
	if err := helper.ValidateUpdatePayment(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// mapper (request → biz)
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

	// mapper (model → response)
	return mapper.BizToPaymentResponse(updated), nil
}

func (uc *PaymentUsecase) verifyCardPin(
	ctx context.Context,
	accountID int64,
	pin string,
) error {

	resp, err := uc.cards.VerifyCardPin(ctx, &cardspb.VerifyCardPinRequest{
		AccountId: accountID,
		Pin:       pin,
	})
	if err != nil {
		return status.Error(codes.Unavailable, "card service unavailable")
	}

	log.Printf(
		"[VerifyCardPin] account_id=%d input_pin=%s verified=%v card_status=%s pin_attempts=%d",
		accountID,
		pin,
		resp.Verified,
		resp.CardStatus,
		resp.PinAttempts,
	)

	if resp.CardStatus == "blocked" {
		return status.Error(codes.PermissionDenied, "card blocked")
	}

	if !resp.Verified {
		return status.Error(codes.PermissionDenied, "invalid pin")
	}

	return nil
}
*/
