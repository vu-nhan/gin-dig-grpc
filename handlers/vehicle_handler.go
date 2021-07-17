package handlers

import (
	"context"
	pb "github.com/vu-nhan/gin-dig-grpc/pb/generated"
	"github.com/vu-nhan/gin-dig-grpc/services"
)

type VehicleHandler struct {
	pb.UnimplementedVehicleServer
	vehicleService services.VehicleService
}

func NewVehicleHandler(injectedVehicleService services.VehicleService) VehicleHandler {
	return VehicleHandler{
		vehicleService: injectedVehicleService,
	}
}

func (h *VehicleHandler) GetAllVehicles(ctx context.Context, request *pb.GetAllVehiclesRequest) (*pb.GetAllVehiclesResponse, error) {
	message, err := h.vehicleService.GetAll()

	if err != nil {
		return nil, err
	}

	return &pb.GetAllVehiclesResponse{
		Message: message,
	}, nil
}

func (h *VehicleHandler) GetVehicleByCode(ctx context.Context, request *pb.GetVehicleByCodeRequest) (*pb.GetVehicleByCodeResponse, error) {
	message, err := h.vehicleService.GetByCode(request.Code)

	if err != nil {
		return nil, err
	}

	return &pb.GetVehicleByCodeResponse{
		Message: message,
	}, nil
}
