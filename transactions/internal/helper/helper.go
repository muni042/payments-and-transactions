package helper

import (
	"errors"
	"transactions/internal/model"
)

func ValidateTransaction(tx *model.Transaction) error {

	if tx.AccountID <= 0 {
		return errors.New("invalid account_id")
	}

	if tx.TransactionType == "" {
		return errors.New("transaction_type is required")
	}
	
	if tx.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if len(tx.Currency) != 3 {
		return errors.New("currency must be 3 characters")
	}

	if tx.Status == "" {
		return errors.New("status is required")
	}

	if tx.TransactionDate.IsZero() {
		return errors.New("transaction_date is required")
	}

	// posting_date is optional, but if present must be valid
	if !tx.PostingDate.IsZero() {
		// already parsed → nothing extra
	}

	return nil
}

/*func ValidateGetTransaction(transactionID int64) error {
	if transactionID <= 0 {
		return errors.New("invalid transaction_id")
	}
	return nil
}*/

func ValidateGetTransaction(transactionID string) error {
	if transactionID == "" {
		return errors.New("transaction_id is required")
	}
	return nil
}


