package client

import (
	"context"

	"github.com/Brenopms/grpc-learning/order/pkg/pb"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) (productServiceClient ProductServiceClient, err error) {
	clientConn, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		return
	}

	client := ProductServiceClient{
		Client: pb.NewProductServiceClient(clientConn),
	}

	return client, nil
}

func (productServiceClient *ProductServiceClient) FindOne(productId int32) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return productServiceClient.Client.FindOne(context.Background(), req)
}

func (productServiceClient *ProductServiceClient) DecreaseStock(productId int32, orderId int32) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return productServiceClient.Client.DecreaseStock(context.Background(), req)
}
