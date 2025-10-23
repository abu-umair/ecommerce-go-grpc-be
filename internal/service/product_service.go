package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/product"
	"github.com/google/uuid"
)

type IProductService interface {
	CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error)
	DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error)
}

type productService struct {
	productRepository repository.IProductRepository
}

func (ps *productService) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	//* cek dulu apakah user admin ?
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims.Role != entity.UserRoleAdmin {
		return nil, utils.UnauthenticatedResponse()
	}

	//* cek juga apakah imagenya ada ? (step ini sementara di skip)
	imagePath := filepath.Join("storage", "product", request.ImageFileName)
	_, err = os.Stat(imagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &product.CreateProductResponse{
				Base: utils.BadRequestResponse("File not found"),
			}, nil
		}
		return nil, err
	}

	//* insert ke db
	productEntity := entity.Product{
		Id:            uuid.NewString(),
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		ImageFileName: request.ImageFileName,
		CreatedAt:     time.Now(),
		CreatedBy:     claims.FullName,
	}

	err = ps.productRepository.CreateNewProduct(ctx, &productEntity)
	if err != nil {
		return nil, err
	}

	//* Success
	return &product.CreateProductResponse{
		Base: utils.SuccessResponse("Product created successfully"),
		Id:   productEntity.Id,
	}, nil
}

func (ps *productService) DetailProduct(ctx context.Context, request *product.DetailProductRequest) (*product.DetailProductResponse, error) {
	//* Query ke DB dengan data Id
	productEntity, err := ps.productRepository.GetProductById(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	//* Apabila null, kita return not found
	if productEntity == nil {
		return &product.DetailProductResponse{
			Base: utils.NotFoundResponse("Product not found"),
		}, nil
	}

	//* Kirim response
	return &product.DetailProductResponse{
			Base:        utils.SuccessResponse("Get product detail success"),
			Id:          productEntity.Id,
			Name:        productEntity.Name,
			Description: productEntity.Description,
			Price:       productEntity.Price,
			ImageUrl:    fmt.Sprintf("http://localhost:3000/storage/product/%s", productEntity.ImageFileName),
		},
		nil

}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &productService{
		productRepository: productRepository,
	}
}
