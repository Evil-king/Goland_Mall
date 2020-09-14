package main

import (
	"Game/rpcservices"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GrpcInit() {
	//创建一个grpc服务器对象
	rpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcservices.RegisterInnerGameInfoServiceServer(rpcServer,new(rpcservices.InnerGameInfoInterface))
	rpcServer.Serve(lis)
}
