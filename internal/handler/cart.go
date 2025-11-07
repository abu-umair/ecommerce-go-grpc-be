package handler

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/cart"
)

type cartHandler struct {
	cart.UnimplementedCartServiceServer
}

func (ch *cartHandler) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	validationErrors, err := utils.CheckValidation(request) //?copas dari auth.go function changePassword
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &cart.AddProductToCartResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}
	res, err := ch.cartService.AddProductToCart(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewCartHandler() *cartHandler {
	return &cartHandler{}
}
