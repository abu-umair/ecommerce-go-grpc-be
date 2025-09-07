package handler //*nama package handler

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/auth"
)

type authHandler struct {
	auth.UnimplementedAuthServiceServer

	authService service.IAuthService
}

func (sh *authHandler) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}
	
	if validationErrors != nil {
		return &auth.RegisterResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}
	
	//?proses Register
	res, err := sh.authService.Register(ctx, request)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

//? mengimplementasikan auth service login
func (sh *authHandler) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}
	
	if validationErrors != nil {
		return &auth.LoginResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}
	
	//?proses Login
	res, err := sh.authService.Login(ctx, request)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func NewAuthHandler(authService service.IAuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}
