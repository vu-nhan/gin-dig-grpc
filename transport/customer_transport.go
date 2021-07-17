package transport

import (
	"context"
	"github.com/vu-nhan/gin-dig-grpc/handlers"
	pb "github.com/vu-nhan/gin-dig-grpc/pb/generated"
)

type CustomerGrpcServer struct {
	pb.UnimplementedCustomerServer
	customerHandler handlers.CustomerHandler
}

func NewCustomerGrpcServer(injectedCustomerHandler handlers.CustomerHandler) *CustomerGrpcServer {
	return &CustomerGrpcServer{
		customerHandler: injectedCustomerHandler,
	}
}

func (s *CustomerGrpcServer) GetAllCustomer(ctx context.Context, request *pb.GetAllCustomerRequest) (*pb.GetAllCustomerResponse, error) {
	return &pb.GetAllCustomerResponse{
		Message: s.customerHandler.GetAllVehicle(),
	}, nil
}

func (s *CustomerGrpcServer) GetCustomerByCode(ctx context.Context, request *pb.GetCustomerByCodeRequest) (*pb.GetCustomerByCodeResponse, error) {
	return &pb.GetCustomerByCodeResponse{
		Message: s.customerHandler.GetVehicleByCode(),
	}, nil
}