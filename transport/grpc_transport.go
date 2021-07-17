package transport

import (
	"github.com/vu-nhan/gin-dig-grpc/handlers"
	pb "github.com/vu-nhan/gin-dig-grpc/pb/generated"
	"github.com/vu-nhan/gin-dig-grpc/provider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GrpcTransportServer struct {
	TransportServer *grpc.Server
}

func NewGrpcTransportServer() *GrpcTransportServer {
	return &GrpcTransportServer{
		TransportServer: grpc.NewServer(),
	}
}

func (s *GrpcTransportServer) StartGrpcTransportServer(listener net.Listener) error {

	container := provider.NewIocContainer()
	container.InitializeDependency()

	if err := container.Invoke(func(vehicleHandler handlers.VehicleHandler) {
		pb.RegisterVehicleServer(s.TransportServer, &vehicleHandler)
	}); err != nil {
		return err
	}

	if err := container.Invoke(func(customerHandler handlers.CustomerHandler) {
		pb.RegisterCustomerServer(s.TransportServer, &customerHandler)
	}); err != nil {
		return err
	}

	reflection.Register(s.TransportServer)

	if err := s.TransportServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
