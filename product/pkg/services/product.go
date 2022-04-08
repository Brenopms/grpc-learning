package services

import (
	"context"
	"net/http"

	"github.com/Brenopms/grpc-learning/pkg/db"
	"github.com/Brenopms/grpc-learning/pkg/models"
	"github.com/Brenopms/grpc-learning/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Stock
	product.Sku = req.Sku

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  []string{result.Error.Error()},
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     int32(product.Id),
	}, nil
}
