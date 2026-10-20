package helper

import (
	"errors"
	"regexp"
	"time"

	pb "sample/api/customers/v1"
)

// CREATE

func ValidateCreateCustomer(req *pb.CreateCustomerRequest) error {

	if req.CustomerNumber == "" {
		return errors.New("customer_number is required")
	}

	if req.FirstName == "" {
		return errors.New("first_name is required")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}

	if req.DateOfBirth != "" {
		if _, err := time.Parse("2006-01-02", req.DateOfBirth); err != nil {
			return errors.New("invalid date_of_birth format (YYYY-MM-DD)")
		}
	}

	if !isValidCustomerStatus(req.Status) {
		return errors.New("invalid status")
	}

	if !isValidKycStatus(req.KycStatus) {
		return errors.New("invalid kyc_status")
	}

	return nil
}

// UPDATE

func ValidateUpdateCustomer(req *pb.UpdateCustomerRequest) error {

	if req.CustomerId <= 0 {
		return errors.New("invalid customer_id")
	}

	if req.Email != "" && !isValidEmail(req.Email) {
		return errors.New("invalid email format")
	}

	if req.Status != "" && !isValidCustomerStatus(req.Status) {
		return errors.New("invalid status")
	}

	if req.KycStatus != "" && !isValidKycStatus(req.KycStatus) {
		return errors.New("invalid kyc_status")
	}

	return nil
}

// DELETE
func ValidateDeleteCustomer(id int64) error {
	if id <= 0 {
		return errors.New("invalid customer_id")
	}
	return nil
}

// GET

func ValidateGetCustomer(id int64) error {
	if id <= 0 {
		return errors.New("invalid customer_id")
	}
	return nil
}

// COMMON

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidCustomerStatus(v string) bool {
	return v == "active" || v == "suspended" || v == "closed"
}

func isValidKycStatus(v string) bool {
	return v == "pending" || v == "verified" || v == "rejected"
}
