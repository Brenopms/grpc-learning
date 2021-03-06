package main

import (
	"log"
	"net"

	"github.com/Brenopms/grpc-learning/product/pkg/config"
	"github.com/Brenopms/grpc-learning/product/pkg/db"
	"github.com/Brenopms/grpc-learning/product/pkg/pb"
	"github.com/Brenopms/grpc-learning/product/pkg/services"
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

	log.Println("Product Service running on port: ", envConfig.Port)

	server := services.Server{
		DbHandler: dbHandler,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalln("Failed to Run grpc server")
	}

}
