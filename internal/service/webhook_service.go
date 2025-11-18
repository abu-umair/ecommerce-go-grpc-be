package service

import (
	"context"
	"errors"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/dto"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
)

type IWebhookService interface {
	ReceiveInvoice(ctx context.Context, request *dto.XenditInvoiceRequest) error
}

type webhookService struct {
	orderRepository repository.IOrderRepository
}

func (ws *webhookService) ReceiveInvoice(ctx context.Context, request *dto.XenditInvoiceRequest) error {
	//* find order di db
	orderEntity, err := ws.orderRepository.GetOrderById(ctx, request.ExternalID)
	if err != nil {
		return err
	}
	if orderEntity == nil {
		return errors.New("order not found")
	}

	//* update entity
	now := time.Now()
	updatedBy := "System"
	orderEntity.OrderStatusCode = entity.OrderStatusCodePaid
	orderEntity.UpdatedAt = &now
	orderEntity.UpdatedBy = &updatedBy
	orderEntity.XenditPaidAt = &now
	orderEntity.XenditPaymentChannel = &request.PaymentChannel
	orderEntity.XenditPaymentMethod = &request.PaymentMethod

	//* proses update db
	err = ws.orderRepository.UpdateOrder(ctx, orderEntity)
	if err != nil {
		return err
	}

	return nil
}

func NewWebhookService(orderRepository repository.IOrderRepository) IWebhookService {
	return &webhookService{
		orderRepository: orderRepository,
	}
}
