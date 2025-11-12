package service

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/pb/order"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
}

type orderService struct{}

func (os *orderService) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	//* simpan 'order' ke database

	//* iterasi semua data product di 'request'
	//* setia iterasinya, simpan 'order_item' ke database
}

func NewOrderService() IOrderService {
	return &orderService{}
}
