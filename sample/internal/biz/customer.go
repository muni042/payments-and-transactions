package biz

import (
	"sample/internal/data"
)

type CustomerUsecase struct {
	repo *data.CustomerRepo
}

func NewCustomerUsecase(repo *data.CustomerRepo) *CustomerUsecase {
	return &CustomerUsecase{repo: repo}
}

/*
// CREATE
func (uc *CustomerUsecase) Create(
	ctx context.Context,
	req *pb.CreateCustomerRequest,
) (*pb.CustomerResponse, error) {

	// validation
	if err := helper.ValidateCreateCustomer(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// map
	c := mapper.CreateCustomerReqToBiz(req)

	// insert
	if err := uc.repo.InsertDB(ctx, c); err != nil {
		return nil, status.Error(codes.Internal, "failed to create customer")
	}

	return mapper.BizToCustomerResponse(c), nil
}

// GET
func (uc *CustomerUsecase) Get(
	ctx context.Context,
	req *pb.GetCustomerRequest,
) (*pb.CustomerResponse, error) {

	if err := helper.ValidateGetCustomer(req.CustomerId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	c, err := uc.repo.GetDB(ctx, req.CustomerId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "customer not found")
	}

	return mapper.BizToCustomerResponse(c), nil
}

// UPDATE
func (uc *CustomerUsecase) Update(
	ctx context.Context,
	req *pb.UpdateCustomerRequest,
) (*pb.CustomerResponse, error) {

	if err := helper.ValidateUpdateCustomer(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	c := mapper.UpdateCustomerReqToBiz(req)

	if err := uc.repo.UpdateDB(ctx, c); err != nil {
		return nil, status.Error(codes.Internal, "customer update failed")
	}

	c, _ = uc.repo.GetDB(ctx, req.CustomerId)
	return mapper.BizToCustomerResponse(c), nil
}

// DELETE
func (uc *CustomerUsecase) Delete(
	ctx context.Context,
	req *pb.DeleteCustomerRequest,
) (*pb.DeleteCustomerResponse, error) {

	if err := helper.ValidateDeleteCustomer(req.CustomerId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := uc.repo.DeleteDB(ctx, req.CustomerId); err != nil {
		return nil, status.Error(codes.Internal, "customer deletion failed")
	}

	return &pb.DeleteCustomerResponse{
		Message: "Customer discarded successfully",
	}, nil
}

// CHECK CUSTOMER EXISTS
func (uc *CustomerUsecase) Exists(
	ctx context.Context,
	req *pb.GetCustomerRequest,
) (*pb.CheckCustomerExistsResponse, error) {

	if err := helper.ValidateGetCustomer(req.CustomerId); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ok, err := uc.repo.ExistsByID(ctx, req.CustomerId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to check customer")
	}

	return &pb.CheckCustomerExistsResponse{Exists: ok}, nil
}
*/
