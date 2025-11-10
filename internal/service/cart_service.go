package service

import (
	"context"
	"fmt"
	"os"
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
	ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error)
	DeleteCart(ctx context.Context, request *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error)
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

func (cs *cartService) ListCart(ctx context.Context, request *cart.ListCartRequest) (*cart.ListCartResponse, error) {
	//* ambil auth user (yang sedang mengakses api ini)
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	//* query list cart dari db
	//* join (table user_cart, product)
	carts, err := cs.cartRepository.GetListCart(ctx, claims.Subject)
	if err != nil {
		return nil, err
	}

	//* build response nya
	var items []*cart.ListCartResponseItem = make([]*cart.ListCartResponseItem, 0)
	for _, cartEntity := range carts {
		item := cart.ListCartResponseItem{
			CartId:          cartEntity.Id,
			ProductId:       cartEntity.Product.Id,
			ProductName:     cartEntity.Product.Name,
			ProductImageUrl: fmt.Sprintf("%s/ product/%s", os.Getenv("STORAGE_SERVICE_URL"), cartEntity.Product.ImageFileName),
			ProductPrice:    cartEntity.Product.Price,
			Quantity:        int64(cartEntity.Quantity),
		}

		items = append(items, &item)
	}

	//* kirim response
	return &cart.ListCartResponse{
		Base:  utils.SuccessResponse("Get list cart success"),
		Items: items,
	}, nil
}

// ?membuat method delete
func (cs *cartService) DeleteCart(ctx context.Context, request *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
	//*mendapatkan user id nya dulu
	//? ambil auth user (yang sedang mengakses api ini)
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	//* mendapatkan data cart
	cartEntity, err := cs.cartRepository.GetCartById(ctx, request.CartId)
	if err != nil {
		return nil, err
	}
	if cartEntity == nil {
		return &cart.DeleteCartResponse{
			Base: utils.NotFoundResponse("Cart not found"),
		}, nil
	}

	//* mencocokan data user id di cart dengan auth (jika tidak cocok , mengirimkan badRequest)
	if cartEntity.UserId != claims.Subject {
		return &cart.DeleteCartResponse{
			Base: utils.BadRequestResponse("Cart user is is not matched"),
		}, nil
	}

	//* delete (hard delete)

	//* kirim response
}

func NewCartService(productRepository repository.IProductRepository, cartRepository repository.ICartRepository) ICartService {
	return &cartService{
		productRepository: productRepository,
		cartRepository:    cartRepository,
	}
}
