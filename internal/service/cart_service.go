package service

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/pb/cart"
)

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
}

type cartService struct{}

func (cs *cartService) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	//* cek terlebi dahulu apakah product id itu ada di db

	//* cek ke db apakah product udah ada di cart user

		//* udah ada -> update

		//* belum -> insert

	//* response
}

func NewCartService() ICartService {
	return &cartService{}
}
