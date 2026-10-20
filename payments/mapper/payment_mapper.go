package mapper

import (
	pb "payments/api/payments/v1"
	"payments/internal/model"
	"strconv"
	"time"
	transactionspb "transactions/api/transactions/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreatePaymentReqToBiz maps CreatePaymentRequest → model.Payment
func CreatePaymentReqToBiz(req *pb.CreatePaymentRequest) (*model.Payment, error) {

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return nil, err
	}

	p := &model.Payment{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		PaymentType:   req.PaymentType,
		PaymentMethod: req.PaymentMethod,
		Amount:        amount,
		Currency:      req.Currency,
		Status:        "initiated",
	}

	if req.ScheduledAt != nil {
		p.ScheduledAt = req.ScheduledAt.AsTime()
	}

	return p, nil
}

// BizToPaymentResponse maps model.Payment → PaymentResponse
func BizToPaymentResponse(p *model.Payment) *pb.PaymentResponse {

	resp := &pb.PaymentResponse{
		PaymentId:     p.PaymentID,
		FromAccountId: p.FromAccountID,
		ToAccountId:   p.ToAccountID,
		PaymentType:   p.PaymentType,
		PaymentMethod: p.PaymentMethod,
		Amount:        strconv.FormatFloat(p.Amount, 'f', 4, 64),
		Currency:      p.Currency,
		Status:        p.Status,
		CreatedAt:     timestamppb.New(p.CreatedAt),
	}

	if !p.ScheduledAt.IsZero() {
		resp.ScheduledAt = timestamppb.New(p.ScheduledAt)
	}

	if !p.ProcessedAt.IsZero() {
		resp.ProcessedAt = timestamppb.New(p.ProcessedAt)
	}

	return resp
}
func PaymentToDebitTransaction(
	fromAccountID int64,
	amount string,
	currency string,
) *transactionspb.AddTransactionRequest {

	return &transactionspb.AddTransactionRequest{
		AccountId:       strconv.FormatInt(fromAccountID, 10),
		TransactionType: transactionspb.TransactionType_TRANSACTION_TYPE_DEBIT,
		Amount:          amount,
		Currency:        currency,
		Status:          transactionspb.TransactionStatus_TRANSACTION_STATUS_POSTED,
		Description:     "Debit for payment",
		PostingDate:     time.Now().Format("2006-01-02"),
		TransactionDate: timestamppb.Now(),
	}
}

// Credit transaction mapper
func PaymentToCreditTransaction(
	toAccountID int64,
	amount string,
	currency string,
) *transactionspb.AddTransactionRequest {

	return &transactionspb.AddTransactionRequest{
		AccountId:       strconv.FormatInt(toAccountID, 10),
		TransactionType: transactionspb.TransactionType_TRANSACTION_TYPE_CREDIT,
		Amount:          amount,
		Currency:        currency,
		Status:          transactionspb.TransactionStatus_TRANSACTION_STATUS_POSTED,
		Description:     "Credit for payment",
		PostingDate:     time.Now().Format("2006-01-02"),
		TransactionDate: timestamppb.Now(),
	}
}
func UpdatePaymentReqToBiz(req *pb.UpdatePaymentRequest) *model.Payment {

	p := &model.Payment{
		PaymentID: req.PaymentId,
	}

	// update status if provided
	if req.Status != "" {
		p.Status = req.Status
	}

	// update processed_at if provided
	if req.ProcessedAt != nil {
		p.ProcessedAt = req.ProcessedAt.AsTime()
	}

	return p
}
