package helper

import (
	pb "accounts/api/accounts/v1"
	"errors"
)

func ValidateCreateAccount(req *pb.CreateAccountRequest) error {

	if req.CustomerId <= 0 {
		return errors.New("customer_id is required")
	}

	if req.AccountNumber == "" {
		return errors.New("account_number is required")
	}

	if req.AccountType == "" {
		return errors.New("account_type is required")
	}

	if req.Currency == "" || len(req.Currency) != 3 {
		return errors.New("currency must be 3 characters like INR")
	}

	if req.Status == "" {
		return errors.New("status is required")
	}

	if req.AvailableBalance < 0 {
		return errors.New("available_balance cannot be negative")
	}

	if req.PendingBalance < 0 {
		return errors.New("pending_balance cannot be negative")
	}

	if req.CreditLimit < 0 {
		return errors.New("credit_limit cannot be negative")
	}

	return nil
}
func ValidateID(id int64, field string) error {
	if id <= 0 {
		return errors.New("invalid " + field)
	}
	return nil
}
