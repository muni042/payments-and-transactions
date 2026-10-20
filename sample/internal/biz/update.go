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
