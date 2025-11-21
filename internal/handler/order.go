package handler

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/order"
)

type orderHandler struct {
	order.UnimplementedOrderServiceServer

	orderService service.IOrderService
}

func (oh *orderHandler) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	//? mengambil validate di func highlightProduct pada product.go
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &order.CreateOrderResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := oh.orderService.CreateOrder(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (oh *orderHandler) ListOrderAdmin(ctx context.Context, request *order.ListOrderAdminRequest) (*order.ListOrderAdminResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &order.ListOrderAdminResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := oh.orderService.ListOrderAdmin(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (oh *orderHandler) ListOrder(ctx context.Context, request *order.ListOrderRequest) (*order.ListOrderResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &order.ListOrderResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := oh.orderService.ListOrder(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewOrderHandler(orderService service.IOrderService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}
