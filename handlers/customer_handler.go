package handlers

import (
	"context"
	pb "github.com/vu-nhan/gin-dig-grpc/pb/generated"
	"github.com/vu-nhan/gin-dig-grpc/services"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServer
	customerService services.CustomerService
}

func NewCustomerHandler(injectedCustomerService services.CustomerService) CustomerHandler {
	return CustomerHandler{
		customerService: injectedCustomerService,
	}
}

func (h *CustomerHandler) GetAllCustomer(ctx context.Context, request *pb.GetAllCustomerRequest) (*pb.GetAllCustomerResponse, error) {
	message, err := h.customerService.GetAll()
	if err != nil {
		return nil, err
	}

	return &pb.GetAllCustomerResponse{
		Message: message,
	}, nil
}

func (h *CustomerHandler) GetCustomerByCode(ctx context.Context, request *pb.GetCustomerByCodeRequest) (*pb.GetCustomerByCodeResponse, error) {
	message, err := h.customerService.GetByCode(request.Code)
	if err != nil {
		return nil, err
	}

	return &pb.GetCustomerByCodeResponse{
		Message: message,
	}, nil
}
