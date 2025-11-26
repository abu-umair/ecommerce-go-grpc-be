package service

import (
	"context"
	"database/sql"
	"fmt"
	operatingsystem "os"
	"runtime/debug"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/order"
	"github.com/google/uuid"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
	ListOrderAdmin(ctx context.Context, request *order.ListOrderAdminRequest) (*order.ListOrderAdminResponse, error)
	ListOrder(ctx context.Context, request *order.ListOrderRequest) (*order.ListOrderResponse, error)
	DetailOrder(ctx context.Context, request *order.DetailOrderRequest) (*order.DetailOrderResponse, error)
}

type orderService struct {
	db                *sql.DB
	orderRepository   repository.IOrderRepository
	productRepository repository.IProductRepository
}

func (os *orderService) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	//? ambil auth user (yang sedang mengakses api ini)
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := os.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if e := recover(); e != nil {
			if tx != nil {
				tx.Rollback() //?rollback jika ada error saan runtime
			}

			debug.PrintStack() //?agar ada stock tracenya yang digunakan utk debug
			panic(e)           //?agar bisa nyampai ke Middleware
		}
	}()

	defer func() {
		if err != nil && tx != nil {
			tx.Rollback() //?rollback jika ada error
		}
	}()

	orderRepo := os.orderRepository.WithTransaction(tx)     //? sydah terintegrasi, dan akan menggantikan semua 'os.orderRepository'
	productRepo := os.productRepository.WithTransaction(tx) //? sydah terintegrasi, dan akan menggantikan semua 'os.productRepository'

	//* simpan 'order' ke database
	numbering, err := orderRepo.GetNumbering(ctx, "order")
	if err != nil {
		return nil, err
	}

	//? mengambil data product dari 'request' untuk menghitung total
	var productIds = make([]string, len(request.Products))
	for i := range request.Products {
		productIds[i] = request.Products[i].Id
	}

	products, err := productRepo.GetProductByIds(ctx, productIds)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]*entity.Product)
	for i := range products {
		productMap[products[i].Id] = products[i]
	}

	var total float64 = 0
	for _, p := range request.Products {
		if productMap[p.Id] == nil { //?menambahkan product notfound
			return &order.CreateOrderResponse{
				Base: utils.NotFoundResponse(fmt.Sprintf("Product %s not found", p.Id)),
			}, nil
		}
		total += productMap[p.Id].Price * float64(p.Quantity)
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
	//? dibuat sebelum diinputkan ke database
	invoiceItems := make([]xendit.InvoiceItem, 0)
	for _, p := range request.Products { //?mapping untuk mengisi invoice product
		prod := productMap[p.Id]
		if prod != nil {
			invoiceItems = append(invoiceItems, xendit.InvoiceItem{
				Name:     prod.Name,
				Price:    prod.Price,
				Quantity: int(p.Quantity),
			})
		}
	}
	xenditInvoice, xenditErr := invoice.CreateWithContext(ctx, &invoice.CreateParams{
		ExternalID: orderEntity.Id,
		Amount:     total,
		Customer: xendit.InvoiceCustomer{
			GivenNames: claims.FullName,
		},
		Currency:           "IDR",
		SuccessRedirectURL: fmt.Sprintf("%s/checkout/%s/success", operatingsystem.Getenv("FRONTEND_BASE_URL"), orderEntity.Id),
		Items:              invoiceItems,
		FailureRedirectURL: fmt.Sprintf("%s/checkout/%s/failure", operatingsystem.Getenv("FRONTEND_BASE_URL"), orderEntity.Id),
	})

	if xenditErr != nil {
		return nil, xenditErr
	}

	orderEntity.XenditInvoiceId = &xenditInvoice.ID
	orderEntity.XenditInvoiceUrl = &xenditInvoice.InvoiceURL

	err = orderRepo.CreateOrder(ctx, &orderEntity)
	if err != nil {
		return nil, err
	}

	// panic("test")

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

		err = orderRepo.CreateOrderItem(ctx, &orderItem)
		if err != nil {
			return nil, err
		}
	}

	numbering.Number++
	err = orderRepo.UpdateNumbering(ctx, numbering)
	if err != nil {
		return nil, err
	}

	err = tx.Commit() //?harus dicommit agar data tersimpan
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		Base: utils.SuccessResponse("Created order success"),
		Id:   orderEntity.Id,
	}, nil
}

func (os *orderService) ListOrderAdmin(ctx context.Context, request *order.ListOrderAdminRequest) (*order.ListOrderAdminResponse, error) {
	//* validasi dulu apakah yang login admin atau bukan
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	//* mengambil query semua order
	orders, metadata, err := os.orderRepository.GetListOrderAdminPagination(ctx, request.Pagination)
	if err != nil {
		return nil, err
	}

	items := make([]*order.ListOrderAdminResponseItem, 0)
	for _, o := range orders {

		products := make([]*order.ListOrderAdminResponseItemProduct, 0)
		for _, oi := range o.Items {
			products = append(products, &order.ListOrderAdminResponseItemProduct{
				Id:       oi.ProductId,
				Name:     oi.ProductName,
				Price:    oi.ProductPrice,
				Quantity: oi.Quantity,
			})
		}

		orderStatusCode := o.OrderStatusCode
		if o.OrderStatusCode == entity.OrderStatusCodeUnpaid && time.Now().After(*o.ExpiredAt) {
			orderStatusCode = entity.OrderStatusCodeExpired
		}

		items = append(items, &order.ListOrderAdminResponseItem{
			Id:         o.Id,
			Number:     o.Number,
			Customer:   o.UserFullName,
			StatusCode: orderStatusCode,
			Total:      o.Total,
			CreatedAt:  timestamppb.New(o.CreatedAt),
			Products:   products,
		})
	}

	return &order.ListOrderAdminResponse{
		Base:       utils.SuccessResponse("List order success"),
		Pagination: metadata,
		Items:      items,
	}, nil
}

func (os *orderService) ListOrder(ctx context.Context, request *order.ListOrderRequest) (*order.ListOrderResponse, error) {
	//?penggunaannya hampir sama dengan ListOrderAdmin, bedanya pada penamaan
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	//* mengambil query order berdasarkan user_idnya
	orders, metadata, err := os.orderRepository.GetListOrderPagination(ctx, request.Pagination, claims.Subject) //?claims.Subject: mengacu user_id nya
	if err != nil {
		return nil, err
	}

	items := make([]*order.ListOrderResponseItem, 0)
	for _, o := range orders {

		products := make([]*order.ListOrderResponseItemProduct, 0)
		for _, oi := range o.Items {
			products = append(products, &order.ListOrderResponseItemProduct{
				Id:       oi.ProductId,
				Name:     oi.ProductName,
				Price:    oi.ProductPrice,
				Quantity: oi.Quantity,
			})
		}

		orderStatusCode := o.OrderStatusCode
		if o.OrderStatusCode == entity.OrderStatusCodeUnpaid && time.Now().After(*o.ExpiredAt) {
			orderStatusCode = entity.OrderStatusCodeExpired
		}

		xenditInvoiceUrl := ""
		if o.XenditInvoiceUrl != nil {
			xenditInvoiceUrl = *o.XenditInvoiceUrl
		}

		items = append(items, &order.ListOrderResponseItem{
			Id:               o.Id,
			Number:           o.Number,
			Customer:         o.UserFullName,
			StatusCode:       orderStatusCode,
			Total:            o.Total,
			CreatedAt:        timestamppb.New(o.CreatedAt),
			Products:         products,
			XenditInvoiceUrl: xenditInvoiceUrl,
		})
	}

	return &order.ListOrderResponse{
		Base:       utils.SuccessResponse("List order success"),
		Pagination: metadata,
		Items:      items,
	}, nil
}

func (os *orderService) DetailOrder(ctx context.Context, request *order.DetailOrderRequest) (*order.DetailOrderResponse, error) {
	//* apakah yang login admin atau bukan
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	//? ambil order berdasarkan requst id
	orderEntity, err := os.orderRepository.GetOrderById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	//? validasi role
	if claims.Role != entity.UserRoleAdmin && claims.Subject != orderEntity.UserId {
		return &order.DetailOrderResponse{
			Base: utils.BadRequestResponse("User id is not mathced"),
		}, nil
	}

	notes := ""
	if orderEntity.Notes != nil {
		notes = *orderEntity.Notes
	}

	XenditInvoiceUrl := ""
	if orderEntity.XenditInvoiceUrl != nil {
		XenditInvoiceUrl = *orderEntity.XenditInvoiceUrl
	}

	// ?ngambil dari func listOrder, diganti variablenya saja
	orderStatusCode := orderEntity.OrderStatusCode
	if orderEntity.OrderStatusCode == entity.OrderStatusCodeUnpaid && time.Now().After(*orderEntity.ExpiredAt) {
		orderStatusCode = entity.OrderStatusCodeExpired
	}

	items := make([]*order.DetailOrderResponseItem, 0)
	for _, oi := range orderEntity.Items {
		items = append(items, &order.DetailOrderResponseItem{
			Id:       oi.ProductId,
			Name:     oi.ProductName,
			Price:    oi.ProductPrice,
			Quantity: oi.Quantity,
		})
	}

	return &order.DetailOrderResponse{
		Base:             utils.SuccessResponse("Get order detail success"),
		Id:               orderEntity.Id,
		Number:           orderEntity.Number,
		UserFullName:     orderEntity.UserFullName,
		Address:          orderEntity.Address,
		PhoneNumber:      orderEntity.PhoneNumber,
		Notes:            notes,
		OrderStatusCode:  orderStatusCode,
		CreatedAt:        timestamppb.New(orderEntity.CreatedAt),
		XenditInvoiceUrl: XenditInvoiceUrl,
		Items:            items,
		Total:            orderEntity.Total,
		ExpiredAt:        timestamppb.New(*orderEntity.ExpiredAt),
	}, nil

}

func NewOrderService(db *sql.DB, orderRepository repository.IOrderRepository, productRepository repository.IProductRepository) IOrderService {
	return &orderService{
		db:                db,
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}
