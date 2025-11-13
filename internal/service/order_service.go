package service

import (
	"context"
	"fmt"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/order"
	"github.com/google/uuid"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
}

type orderService struct {
	orderRepository   repository.IOrderRepository
	productRepository repository.IProductRepository
}

func (os *orderService) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	//* simpan 'order' ke database
	numbering, err := os.orderRepository.GetNumbering(ctx, "order")
	if err != nil {
		return nil, err
	}

	//? mengambil data product dari 'request' untuk menghitung total
	var productIds = make([]string, len(request.Products))
	for i := range request.Products {
		productIds[i] = request.Products[i].Id
	}

	products, err := os.productRepository.GetProductByIds(ctx, productIds)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]*entity.Product)
	for i := range products {
		productMap[products[i].Id] = products[i]
	}

	var total float64 = 0
	for _, p := range request.Products {
		total += productMap[p.Id].Price * float64(p.Quantity)
	}

	//? ambil auth user (yang sedang mengakses api ini)
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expiredAt := now.Add(24 * time.Hour)
	orderEntity := entity.Order{
		Id:              uuid.NewString(),
		Number:          fmt.Sprintf("ORD-%d%08d", now.Year(), numbering.Number), // ORD-202500000010
		UserId:          claims.Subject,
		OrderStatusCode: entity.OrderStatusCodeUnpaid,
		UserFullName:    request.FullName,
		Address:         request.Address,
		PhoneNumber:     request.PhoneNumber,
		Notes:           &request.Notes,
		Total:           total,
		ExpiredAt:       &expiredAt,
		CreatedAt:       now,
		CreatedBy:       claims.FullName,
	}

	err = os.orderRepository.CreateOrder(ctx, &orderEntity)
	if err != nil {
		return nil, err
	}

	//* iterasi semua data product di 'request'
	//* setiap iterasinya, simpan 'order_item' ke database
	for _, p := range request.Products {
		var orderItem = entity.OrderItem{
			Id:                   uuid.NewString(),
			ProductId:            p.Id,
			ProductName:          productMap[p.Id].Name, //?mengambil data dari maping
			ProductImageFileName: productMap[p.Id].ImageFileName,
			ProductPrice:         productMap[p.Id].Price,
			Quantity:             int64(p.Quantity), //?diakses dari requestnya
			OrderId:              orderEntity.Id,
			CreatedAt:            now,
			CreatedBy:            claims.FullName,
		}

		err = os.orderRepository.CreateOrderItem(ctx, &orderItem)
		if err != nil {
			return nil, err
		}
	}

	numbering.Number++
	err = os.orderRepository.UpdateNumbering(ctx, numbering)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{
		Base: utils.SuccessResponse("Created order success"),
		Id:   orderEntity.Id,
	}, nil
}

func NewOrderService(orderRepository repository.IOrderRepository, productRepository repository.IProductRepository) IOrderService {
	return &orderService{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}
