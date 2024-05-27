package main

import (
	"context"
	"log"
	"net"

	"github.com/brunodev09/industrial-oms-go/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.Env("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcSrv := grpc.NewServer()

	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	store := NewStore()
	service := NewService(store)

	NewGRPCHandler(grpcSrv)
	log.Println("GRPC Server has been started at ", grpcAddr)

	service.CreateOrder(context.Background())

	if err := grpcSrv.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
