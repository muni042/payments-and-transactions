package service

import (
	"context"

	pb "sample/api/customers/v1"
	"sample/internal/biz"
)

type CustomersServiceService struct {
	pb.UnimplementedCustomersServiceServer
	uc *biz.CustomerUsecase
}

func NewCustomersServiceService(uc *biz.CustomerUsecase) *CustomersServiceService {
	return &CustomersServiceService{uc: uc}
}

func (s *CustomersServiceService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CustomerResponse, error) {
	return s.uc.Create(ctx, req)
}

func (s *CustomersServiceService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.CustomerResponse, error) {
	return s.uc.Get(ctx, req)
}

func (s *CustomersServiceService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.CustomerResponse, error) {
	return s.uc.Update(ctx, req)
}

func (s *CustomersServiceService) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	return s.uc.Delete(ctx, req)
}

// 🔥 This is the previously unimplemented function
func (s *CustomersServiceService) CheckCustomerExists(ctx context.Context, req *pb.GetCustomerRequest) (*pb.CheckCustomerExistsResponse, error) {

	_, err := s.uc.Get(ctx, req)
	if err != nil {
		return &pb.CheckCustomerExistsResponse{Exists: false}, nil
	}

	return &pb.CheckCustomerExistsResponse{Exists: true}, nil
}
