package helper

import (
	"errors"
	"strconv"
	"strings"
	"time"

	pb "cards/api/cards/v1"
)

var validCardTypes = map[string]bool{
	"debit":   true,
	"credit":  true,
	"prepaid": true,
}

var validCardStatus = map[string]bool{
	"active":  true,
	"blocked": true,
	"expired": true,
}

// Validate Create
func ValidateCreateCard(req *pb.CreateCardRequest) error {

	if req.AccountId <= 0 {
		return errors.New("account_id must be positive")
	}

	if len(strings.TrimSpace(req.CardNumberHash)) < 20 {
		return errors.New("invalid card_number_hash")
	}

	if !validCardTypes[strings.ToLower(req.CardType)] {
		return errors.New("invalid card_type")
	}

	expiry, err := time.Parse("2006-01-02", req.ExpiryDate)
	if err != nil || expiry.Before(time.Now()) {
		return errors.New("invalid or expired expiry_date")
	}

	daily, err := strconv.ParseFloat(req.DailyLimit, 64)
	if err != nil || daily < 0 {
		return errors.New("invalid daily_limit")
	}

	monthly, err := strconv.ParseFloat(req.MonthlyLimit, 64)
	if err != nil || monthly < daily {
		return errors.New("monthly_limit must be >= daily_limit")
	}
	/*if pin=""{
		return errors.New("give pin value")
	}*/

	return nil
}

// Validate Update
func ValidateUpdateCard(req *pb.UpdateCardRequest) error {

	if req.CardId <= 0 {
		return errors.New("invalid card_id")
	}

	if !validCardStatus[strings.ToLower(req.CardStatus)] {
		return errors.New("invalid card_status")
	}

	daily, err := strconv.ParseFloat(req.DailyLimit, 64)
	if err != nil || daily < 0 {
		return errors.New("invalid daily_limit")
	}

	monthly, err := strconv.ParseFloat(req.MonthlyLimit, 64)
	if err != nil || monthly < daily {
		return errors.New("monthly_limit must be >= daily_limit")
	}

	if req.PinAttempts < 0 || req.PinAttempts > 3 {
		return errors.New("pin_attempts must be between 0 and 3")
	}

	return nil
}

// Validate ID
func ValidateCardID(id int64) error {
	if id <= 0 {
		return errors.New("invalid card_id")
	}
	return nil
}
