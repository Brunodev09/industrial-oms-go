package main

import (
	"context"
	"log"

	pb "github.com/brunodev09/industrial-oms-go/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("=> A new order has been placed!")
	order := &pb.Order{
		Id: "322",
	}
	return order, nil
}
