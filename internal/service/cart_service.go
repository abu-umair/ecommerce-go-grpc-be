package service

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/pb/cart"
)

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
}

type cartService struct{}

func(cs *cartService) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	
}
func NewCartService() ICartService {
	return &cartService{}
}