package grpcmiddleware

import (
	"context"

	jwtentity"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	gocache "github.com/patrickmn/go-cache"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type authMiddleware struct {
	cacheService *gocache.Cache
}

func (am *authMiddleware) Middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// Ambil token  dari metadata
	tokenStr, err := jwt.ParseTokenFromContext(ctx)
	if err != nil {
	    return nil, err
	}

	// Cek token  dari  logout cache

	// Parse jwt nya hingga jadi entity

	// Sematkan entity ke context

	res, err := handler(ctx, req)

	return res, err
}

func NewAuthMiddleware(cacheService *gocache.Cache) *authMiddleware {
    return &authMiddleware{
        cacheService: cacheService,
    }
}
