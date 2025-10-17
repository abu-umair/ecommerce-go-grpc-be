package handler //*nama package handler

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/product"
)

type productHandler struct {
	product.UnimplementedProductServiceServer

	authService service.IAuthService
}

func (ph *productHandler) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &product.CreateProductResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	//?proses Register
	res, err := ph.authService.Register(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewProductHandler(authService service.IAuthService) *productHandler {
	return &productHandler{
		authService: authService,
	}
}
