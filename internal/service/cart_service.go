package service

import (
	"context"

	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/cart"
)

type ICartService interface {
	AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error)
}

type cartService struct {
	productRepository repository.IProductRepository
	cartRepository    repository.ICartRepository
}

func (cs *cartService) AddProductToCart(ctx context.Context, request *cart.AddProductToCartRequest) (*cart.AddProductToCartResponse, error) {
	//? mengambil claimsnya user (siapapun dapat mengakses ini)
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	//* cek terlebi dahulu apakah product id itu ada di db
	productEntity, err := cs.productRepository.GetProductById(ctx, request.ProductId)
	if err != nil {
		return nil, err
	}
	if productEntity == nil {
		return &cart.AddProductToCartResponse{
			Base: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	//* cek ke db apakah product udah ada di cart user
	cartEntity, err := cs.cartRepository.GetCartByProductAndUserId(ctx, request.ProductId, claims.Subject)
	if err != nil {
		return nil, err
	}

	//** udah ada -> update
	if cartEntity != nil {
		return &cart.AddProductToCartResponse{
			Base: utils.SuccessResponse("Add Product to cart Success"),
			Id:   "",
		}, nil
	}

	//** belum -> insert

	//* response
	return &cart.AddProductToCartResponse{
		Base: utils.SuccessResponse("Add Product to cart Success"),
		Id:   "",
	}, nil
}

func NewCartService(productRepository repository.IProductRepository, cartRepository repository.ICartRepository) ICartService {
	return &cartService{
		productRepository: productRepository,
		cartRepository:    cartRepository,
	}
}
