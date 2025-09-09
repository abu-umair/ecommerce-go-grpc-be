package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
	Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error)
	Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error)
}

type authService struct {
	authRepository repository.IAuthRepository
}

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if request.Password != request.PasswordConfirmation {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("Password is not matched"),
		}, nil
	}
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)
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

	err = as.authRepository.InsertUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Base: utils.SuccessResponse("User is registered"),
	}, nil
}

func (as *authService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	//? implement logic dari loginnya
	//* check apakah email ada di database
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return &auth.LoginResponse{
			Base: utils.BadRequestResponse("User is not registered"),
		}, nil
	}

	//* check apakah password sama dengan password di database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
		}
		return nil, err
	}

	//* generate jwt
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Id,
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.RoleCode,
	})
	secretKey := os.Getenv("JWT_SECRET")
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	//* kirim response
	return &auth.LoginResponse{
		Base:        utils.SuccessResponse("Login successful"),
		AccessToken: accessToken,
	}, nil

}

// Logout implements
func (as *authService) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	//? dapatkan token dari metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	bearerToken, ok := md["authorization"]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	if len(bearerToken) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	// Bearer qkfqkfqkqwkfqwkq...
	tokenSplit := strings.Split(bearerToken[0], " ")

	if len(tokenSplit) != 2 {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	if tokenSplit[0] != "Bearer" {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	jwtToken := tokenSplit[1]

	//? kembalikan token tadi hingga menjadi entity jwt
	

	//? kita masukkan token ke dalam memory db / cache

	//? kirim response
}

func NewAuthService(authRepository repository.IAuthRepository) IAuthService {
	return &authService{
		authRepository: authRepository,
	}
}
