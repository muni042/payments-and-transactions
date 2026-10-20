package mapper

import (
	pb "sample/api/customers/v1"
	"sample/internal/model"

	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateCustomerReqToBiz(req *pb.CreateCustomerRequest) *model.Customer {

	first := req.FirstName
	last := req.LastName
	email := req.Email
	phone := req.Phone
	status := req.Status
	kyc := req.KycStatus

	var dob *time.Time
	if req.DateOfBirth != "" {
		t, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err == nil {
			dob = &t
		}
	}

	return &model.Customer{
		CustomerNumber: req.CustomerNumber,
		FirstName:      &first,
		LastName:       &last,
		Email:          &email,
		Phone:          &phone,
		DateOfBirth:    dob,
		Status:         &status,
		KycStatus:      &kyc,
	}
}

func UpdateCustomerReqToBiz(req *pb.UpdateCustomerRequest) *model.Customer {
	return &model.Customer{
		CustomerID: req.CustomerId,
		FirstName:  &req.FirstName,
		LastName:   &req.LastName,
		Email:      &req.Email,
		Phone:      &req.Phone,
		Status:     &req.Status,
		KycStatus:  &req.KycStatus,
	}
}

func BizToCustomerResponse(c *model.Customer) *pb.CustomerResponse {

	res := &pb.CustomerResponse{
		CustomerId:     c.CustomerID,
		CustomerNumber: c.CustomerNumber,
		FirstName:      deref(c.FirstName),
		LastName:       deref(c.LastName),
		Email:          deref(c.Email),
		Phone:          deref(c.Phone),
		Status:         deref(c.Status),
		KycStatus:      deref(c.KycStatus),
	}

	if c.CreationTime != nil {
		res.CreatedAt = timestamppb.New(*c.CreationTime)
	}
	if c.ModifiedTime != nil {
		res.UpdatedAt = timestamppb.New(*c.ModifiedTime)
	}

	return res
}
func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
