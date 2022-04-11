package main

import (
	"log"
	"net"

	"github.com/Brenopms/grpc-learning/order/pkg/client"
	"github.com/Brenopms/grpc-learning/order/pkg/config"
	"github.com/Brenopms/grpc-learning/order/pkg/db"
	"github.com/Brenopms/grpc-learning/order/pkg/pb"
	service "github.com/Brenopms/grpc-learning/order/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	envConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Could not load environment configuration: ", err)
	}

	dbHandler := db.Init(envConfig.DBUrl)

	netListener, err := net.Listen("tcp", envConfig.Port)

	if err != nil {
		log.Fatalln("Failed to Start TCP Server")
	}

	log.Println("Product Service running on port:", envConfig.Port)

	productService, err := client.InitProductServiceClient(envConfig.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed at starting Product Service Client:", err)
	}

	server := service.Server{
		DbHandler:            dbHandler,
		ProductServiceClient: productService,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalln("Failed to start the server:", err)
	}
}
