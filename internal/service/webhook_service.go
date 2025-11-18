package service

import (
	"context"

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

	//* update entity

	//* proses update db
	return nil
}

func NewWebhookService(orderRepository repository.IOrderRepository) IWebhookService {
	return &webhookService{
		orderRepositoru: orderRepository,
	}
}
