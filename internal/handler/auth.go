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

// ? mengimplementasikan auth service login
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

// ? mengimplementasikan auth service logout
func (sh *authHandler) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &auth.LogoutResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	//?proses Logout
	res, err := sh.authService.Logout(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//? mengimplementasikan auth service change password
func (sh *authHandler) ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &auth.ChangePasswordResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	//?proses ChangePassword
	res, err := sh.authService.ChangePassword(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
} 

//? mengimplementasikan get profile
func (sh *authHandler) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	//?proses 
	res, err := sh.authService.GetProfile(ctx, request)
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
