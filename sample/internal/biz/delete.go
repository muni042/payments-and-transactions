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