package main

import (
	"fmt"
	"github.com/vu-nhan/gin-dig-grpc/handlers"
	pb "github.com/vu-nhan/gin-dig-grpc/pb/generated"
	"github.com/vu-nhan/gin-dig-grpc/provider"
	"github.com/vu-nhan/gin-dig-grpc/transport"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main()  {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}


	container := provider.NewIocContainer()
	container.InitializeDependency()

	grpcServer := grpc.NewServer()
	if err := container.Invoke(func(vehicleHandler handlers.VehicleHandler) {
		vehicleGrpcServer := transport.NewGrpcServer(vehicleHandler)
		pb.RegisterVehicleServer(grpcServer, vehicleGrpcServer)
	}); err != nil {
		fmt.Printf("Invoke Vehicle Handler error \n %s", err.Error())
	}

	if err := container.Invoke(func(customerHandler handlers.CustomerHandler) {
		customerGrpcServer := transport.NewCustomerGrpcServer(customerHandler)
		pb.RegisterCustomerServer(grpcServer, customerGrpcServer)
	}); err != nil {
		fmt.Printf("Invoke Customer Handler error \n %s", err.Error())
	}

	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}