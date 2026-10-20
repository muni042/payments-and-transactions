package helper

import (
	"errors"
	pb "payments/api/payments/v1"
	"strconv"
)

// ValidateCreatePayment validates CreatePaymentRequest
func ValidateCreatePayment(req *pb.CreatePaymentRequest) error {

	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.FromAccountId <= 0 {
		return errors.New("from_account_id must be greater than zero")
	}

	if req.ToAccountId <= 0 {
		return errors.New("to_account_id must be greater than zero")
	}

	if req.FromAccountId == req.ToAccountId {
		return errors.New("from_account_id and to_account_id cannot be the same")
	}

	if req.PaymentType == "" {
		return errors.New("payment_type is required")
	}

	if req.PaymentMethod == "" {
		return errors.New("payment_method is required")
	}

	if req.Currency == "" {
		return errors.New("currency is required")
	}

	if req.Amount == "" {
		return errors.New("amount is required")
	}

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return errors.New("amount must be a valid number")
	}

	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return nil
}
func ValidateGetPayment(req *pb.GetPaymentRequest) error {
	if req.PaymentId <= 0 {
		return errors.New("invalid payment_id")
	}
	return nil
}
func ValidateUpdatePayment(req *pb.UpdatePaymentRequest) error {
	if req.PaymentId <= 0 {
		return errors.New("invalid payment_id")
	}

	// at least one updatable field must be present
	if req.Status == "" && req.ProcessedAt == nil {
		return errors.New("nothing to update")
	}

	return nil
}
