package transport

import (
	"context"
	"github.com/vu-nhan/gin-dig-grpc/handlers"
	pb "github.com/vu-nhan/gin-dig-grpc/pb/generated"
)

type VehicleGrpcServer struct{
	pb.UnimplementedVehicleServer
	vehicleHandler handlers.VehicleHandler
}

func NewGrpcServer(injectedVehicleHandler handlers.VehicleHandler) *VehicleGrpcServer {
	return &VehicleGrpcServer{
		vehicleHandler:             handlers.VehicleHandler{},
	}
}

func (s *VehicleGrpcServer) GetAll(ctx context.Context, request *pb.GetAllVehiclesRequest) (*pb.GetAllVehiclesResponse, error) {
	return &pb.GetAllVehiclesResponse{
		Message: s.vehicleHandler.GetAllVehicle(),
	}, nil
}

func (s *VehicleGrpcServer) GetByCode(ctx context.Context, request *pb.GetVehicleByCodeRequest) (*pb.GetVehicleByCodeResponse, error) {
	return &pb.GetVehicleByCodeResponse{
		Message: s.vehicleHandler.GetVehicleByCode(),
	}, nil
}