// CHECK CUSTOMER EXISTS
package biz
import (
	"context"

	pb "sample/api/customers/v1"
	//"sample/internal/data"
	"sample/internal/helper"
	//"sample/internal/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
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
