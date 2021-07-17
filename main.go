package main

import (
	"github.com/vu-nhan/gin-dig-grpc/transport"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcTransport := transport.NewGrpcTransportServer()
	if err := grpcTransport.StartGrpcTransportServer(listener); err != nil {
		println(err.Error())
		panic(err)
	}


}
