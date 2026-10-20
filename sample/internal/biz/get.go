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