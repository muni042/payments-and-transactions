package biz

import (
	customerspb "sample/api/customers/v1"

	"accounts/internal/data"
)

type AccountUsecase struct {
	repo *data.AccountRepo
	cust customerspb.CustomersServiceClient
}

func NewAccountUsecase(
	repo *data.AccountRepo,
	cust customerspb.CustomersServiceClient,
) *AccountUsecase {
	return &AccountUsecase{
		repo: repo,
		cust: cust,
	}
}

/*
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

// GET ACCOUNT
func (as *AccountUsecase) Get(
	ctx context.Context,
	req *pb.GetAccountRequest,
) (*pb.AccountResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	account, err := as.repo.GetDB(ctx, req.AccountId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	return mapper.BizToAccountResponse(account), nil
}

// UPDATE ACCOUNT
func (as *AccountUsecase) Update(
	ctx context.Context,
	req *pb.UpdateAccountRequest,
) (*pb.AccountResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	account := mapper.UpdateReqtoBiz(req)

	if err := as.repo.UpdateDB(ctx, account); err != nil {
		return nil, status.Error(codes.Internal, "account update failed")
	}

	account, _ = as.repo.GetDB(ctx, req.AccountId)
	return mapper.BizToAccountResponse(account), nil
}

// DELETE ACCOUNT
func (as *AccountUsecase) Delete(
	ctx context.Context,
	req *pb.DeleteAccountRequest,
) (*pb.DeleteAccountResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := as.repo.DeleteDB(ctx, req.AccountId); err != nil {
		return nil, status.Error(codes.Internal, "account deletion failed")
	}

	return &pb.DeleteAccountResponse{
		Message: "Account discarded successfully",
	}, nil
}

// CREDIT
func (as *AccountUsecase) Credit(
	ctx context.Context,
	req *pb.CreditAccountRequest,
) (*pb.OperationResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.Amount <= 0 {
		return nil, status.Error(codes.InvalidArgument, "amount must be greater than zero")
	}

	if err := as.repo.Credit(ctx, req.AccountId, float64(req.Amount)); err != nil {
		return nil, status.Error(codes.Internal, "credit failed")
	}

	return &pb.OperationResponse{Success: true}, nil
}

// DEBIT
func (as *AccountUsecase) Debit(
	ctx context.Context,
	req *pb.DebitAccountRequest,
) (*pb.OperationResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.Amount <= 0 {
		return nil, status.Error(codes.InvalidArgument, "amount must be greater than zero")
	}

	ok, err := as.repo.Debit(ctx, req.AccountId, float64(req.Amount))
	if err != nil {
		return nil, status.Error(codes.Internal, "debit failed")
	}

	if !ok {
		return nil, status.Error(codes.FailedPrecondition, "insufficient balance")
	}

	return &pb.OperationResponse{Success: true}, nil
}

// CHECK ACCOUNT EXISTS
func (as *AccountUsecase) Exists(
	ctx context.Context,
	req *pb.GetAccountRequest,
) (*pb.CheckAccountExistsResponse, error) {

	if err := helper.ValidateID(req.AccountId, "account_id"); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ok, err := as.repo.ExistsByID(ctx, req.AccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to check account")
	}

	return &pb.CheckAccountExistsResponse{Exists: ok}, nil
}
*/
