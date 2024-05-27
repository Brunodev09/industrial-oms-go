package main

import (
	"log"
	"net/http"

	common "github.com/brunodev09/industrial-oms-go/common"
	pb "github.com/brunodev09/industrial-oms-go/common/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr         = common.Env("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {

	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at the address ", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("HTTP Server listening at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal(err)
	}
}
