package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/product"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/product"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IProductService interface {
	CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error)
}

type productService struct {
	authRepository repository.IAuthRepository
	cacheService   *gocache.Cache
}

func (ps *productService) CreateProduct(ctx context.Context, request *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	if request.Password != request.PasswordConfirmation {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("Password is not matched"),
		}, nil
	}
	user, err := ps.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("User already exist"),
		}, nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}

	// Insert ke db
	newUser := entity.User{
		Id:        uuid.NewString(),
		FullName:  request.FullName,
		Email:     request.Email,
		Password:  string(hashedPassword),
		RoleCode:  entity.UserRoleCustomer,
		CreatedAt: time.Now(),
		CreatedBy: &request.FullName,
	}

	err = ps.authRepository.InsertUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &product.CreateProductResponse{
		Base: utils.SuccessResponse("User is registered"),
	}, nil
}


func NewProductService(authRepository repository.IAuthRepository, cacheService *gocache.Cache) IProductService {
	return &productService{
		authRepository: authRepository,
		cacheService:   cacheService,
	}
}
