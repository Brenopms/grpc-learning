package services

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/Brenopms/grpc-learning/pkg/db"
	"github.com/Brenopms/grpc-learning/pkg/models"
	"github.com/Brenopms/grpc-learning/pkg/pb"
)

type Server struct {
	DbHandler db.Handler
	pb.UnimplementedProductServiceServer
}

var validate *validator.Validate

func getValidationErrors(err error) []string {
	requestErrors := []string{}
	validationErrors := err.(validator.ValidationErrors)
	for _, validationError := range validationErrors {
		requestErrors = append(requestErrors, validationError.Error())
	}

	return requestErrors
}

func (server *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price
	product.Sku = req.Sku

	validate = validator.New()
	err := validate.Struct(product)

	if err != nil {
		validationErrors := getValidationErrors(err)

		return &pb.CreateProductResponse{
			Status: http.StatusBadRequest,
			Error:  validationErrors,
		}, nil
	}

	if result := server.DbHandler.DB.Create(&product); result.Error != nil {
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

func (server *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if result := server.DbHandler.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  []string{result.Error.Error()},
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (server *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product
	var decreaseStockLog models.StockDecreaseLog

	if result := server.DbHandler.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  []string{result.Error.Error()},
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  []string{"stock too low"},
		}, nil
	}

	// Check if the order request was already processed and already decreased a quantity
	if result := server.DbHandler.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&decreaseStockLog); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  []string{"Stock already decreased"},
		}, nil
	}

	product.Stock = product.Stock - 1

	server.DbHandler.DB.Save(&product)

	decreaseStockLog.OrderId = req.OrderId
	decreaseStockLog.ProductId = product.Id

	server.DbHandler.DB.Create(&decreaseStockLog)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil

}
