package biz

import (
	"context"

	pb "accounts/api/accounts/v1"
	customerspb "sample/api/customers/v1"

	"accounts/internal/helper"
	"accounts/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CREATE ACCOUNT
func (as *AccountUsecase) Create(
	ctx context.Context,
	req *pb.CreateAccountRequest,
) (*pb.AccountResponse, error) {

	//validation (helper)
	if err := helper.ValidateCreateAccount(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//check customer exists
	existsResp, err := as.cust.CheckCustomerExists(
		ctx,
		&customerspb.GetCustomerRequest{CustomerId: req.CustomerId},
	)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "customer service unavailable")
	}
	if !existsResp.Exists {
		return nil, status.Error(codes.NotFound, "customer does not exist")
	}

	// map request → model
	account := mapper.CreateReqtoBiz(req)

	//  save
	if err := as.repo.InsertDB(ctx, account); err != nil {
		return nil, status.Error(codes.Internal, "failed to create account")
	}

	// response
	return mapper.BizToAccountResponse(account), nil
}
