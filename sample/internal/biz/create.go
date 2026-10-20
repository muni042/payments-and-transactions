package biz
import (
	"context"

	pb "sample/api/customers/v1"
	//"sample/internal/data"
	"sample/internal/helper"
	"sample/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
// CREATE
func (uc *CustomerUsecase) Create(
	ctx context.Context,
	req *pb.CreateCustomerRequest,
) (*pb.CustomerResponse, error) {

	// 1️⃣ validation
	if err := helper.ValidateCreateCustomer(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 2️⃣ map
	c := mapper.CreateCustomerReqToBiz(req)

	// 3️⃣ insert
	if err := uc.repo.InsertDB(ctx, c); err != nil {
		return nil, status.Error(codes.Internal, "failed to create customer")
	}

	return mapper.BizToCustomerResponse(c), nil
}