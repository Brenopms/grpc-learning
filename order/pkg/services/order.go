package service

import (
	"context"
	"log"
	"net/http"

	"github.com/Brenopms/grpc-learning/order/pkg/client"
	"github.com/Brenopms/grpc-learning/order/pkg/db"
	"github.com/Brenopms/grpc-learning/order/pkg/models"
	"github.com/Brenopms/grpc-learning/order/pkg/pb"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	DbHandler            db.Handler
	ProductServiceClient client.ProductServiceClient
	pb.UnimplementedOrderServiceServer
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

func calculateOrderPrice(productPrice float64, orderQuantity int32) float64 {
	return productPrice * float64(orderQuantity)
}

func (server *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Checking if product id was not sent in the request
	if req.ProductId == 0 {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: []string{"Invalid Product Id Requested"}}, nil
	}
	product, err := server.ProductServiceClient.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: []string{err.Error()}}, nil
	} else if product.Status == http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: []string{"Not enough stock for the requested product"}}, nil
	}

	order := models.Order{
		Price:     calculateOrderPrice(product.Data.Price, req.Quantity),
		ProductId: product.Data.Id,
		UserId:    req.UserId,
		Quantity:  req.Quantity,
	}

	validate = validator.New()
	err = validate.Struct(order)

	if err != nil {
		validationErrors := getValidationErrors(err)

		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  validationErrors,
		}, nil
	}

	server.DbHandler.DB.Create(&order)

	res, err := server.ProductServiceClient.DecreaseStock(req.ProductId, order.Id, order.Quantity)
	log.Println(err, res)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: []string{err.Error()}}, nil
	} else if res.Status == http.StatusConflict || res.Status == http.StatusBadRequest {
		// Order was already placed but not computed
		server.DbHandler.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{Status: res.Status, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
