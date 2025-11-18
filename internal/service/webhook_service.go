package service

import (
	"context"
	"errors"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/dto"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
)

type IWebhookService interface {
	ReceiveInvoice(ctx context.Context, request *dto.XenditInvoiceRequest) error
}

type webhookService struct {
	orderRepositoru repository.IOrderRepository
}

func (ws *webhookService) ReceiveInvoice(ctx context.Context, request *dto.XenditInvoiceRequest) error {
	//* find order di db
	orderEntity, err := ws.orderRepositoru.GetOrderById(ctx, request.ExternalID)
	if err != nil {
		return err
	}
	if orderEntity == nil {
		return errors.New("order not found")
	}

	//* update entity
	now := time.Now()
	updatedBy := "System"
	orderEntity.UpdatedAt = &now
	orderEntity.UpdatedBy = &updatedBy
	orderEntity.XenditPaidAt = &now
	orderEntity.XenditPaymentChannel = &request.PaymentChannel
	orderEntity.XenditPaymentMethod = &request.PaymentMethod

	//* proses update db
	return nil
}

func NewWebhookService(orderRepository repository.IOrderRepository) IWebhookService {
	return &webhookService{
		orderRepositoru: orderRepository,
	}
}
