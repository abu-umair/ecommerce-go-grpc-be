package service

import (
	"context"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/cart"
	"github.com/google/uuid"
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
		now := time.Now()
		cartEntity.Quantity += 1
		cartEntity.UpdatedAt = &now
		cartEntity.UpdatedBy = &claims.Subject

		err = cs.cartRepository.UpdateCart(ctx, cartEntity)
		if err != nil {
			return nil, err
		}

		return &cart.AddProductToCartResponse{
			Base: utils.SuccessResponse("Add Product to cart Success"),
			Id:   cartEntity.Id,
		}, nil
	}

	//** belum -> insert/create
	newCartEntity := entity.UserCart{
		Id:        uuid.NewString(),
		UserId:    claims.Subject,
		ProductId: request.ProductId,
		Quantity:  1,
		CreatedAt: time.Now(),
		CreatedBy: claims.FullName,
	}

	err = cs.cartRepository.CreateNewCart(ctx, &newCartEntity)
	if err != nil {
		return nil, err
	}

	//* response
	return &cart.AddProductToCartResponse{
		Base: utils.SuccessResponse("Add Product to cart Success"),
		Id:   newCartEntity.Id,
	}, nil
}

func NewCartService(productRepository repository.IProductRepository, cartRepository repository.ICartRepository) ICartService {
	return &cartService{
		productRepository: productRepository,
		cartRepository:    cartRepository,
	}
}
