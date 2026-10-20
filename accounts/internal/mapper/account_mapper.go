package mapper

import (
	pb "accounts/api/accounts/v1"
	"accounts/internal/model"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateReqtoBiz(req *pb.CreateAccountRequest) *model.Account {
	a := model.Account{
		CustomerID:       req.CustomerId,
		AccountNumber:    req.AccountNumber,
		AccountType:      req.AccountType,
		Currency:         req.Currency,
		Status:           req.Status,
		AvailableBalance: float64(req.AvailableBalance),
	}
	return &a
}
func UpdateReqtoBiz(req *pb.UpdateAccountRequest) *model.Account {
	a := model.Account{
		AccountID:        req.AccountId,
		AccountNumber:    req.AccountNumber,
		AccountType:      req.AccountType,
		Currency:         req.Currency,
		Status:           req.Status,
		AvailableBalance: float64(req.AvailableBalance),
	}
	return &a
}
func BizToAccountResponse(a *model.Account) *pb.AccountResponse {
	res := pb.AccountResponse{
		AccountId:         a.AccountID,
		AccountNumber:     a.AccountNumber,
		Currency:          a.Currency,
		Status:            a.Status,
		AvailableBalance:  strconv.FormatFloat(a.AvailableBalance, 'f', -1, 64),
		PendingBalance:    strconv.FormatFloat(a.PendingBalance, 'f', -1, 64),
		CreditLimit:       strconv.FormatFloat(a.CreditLimit, 'f', -1, 64),
		CreatedAt:         timestamppb.New(a.CreatedAt),
		UpdatedAt:         timestamppb.New(a.UpdatedAt),
		LastTransactionAt: timestamppb.New(a.LastTransactionAt),
	}
	return &res
}
