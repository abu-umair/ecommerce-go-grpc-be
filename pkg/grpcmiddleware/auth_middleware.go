package grpcmiddleware

import (
	"context"

	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	gocache "github.com/patrickmn/go-cache"

	"google.golang.org/grpc"
)

type authMiddleware struct {
	cacheService *gocache.Cache
}

func (am *authMiddleware) Middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// Ambil token  dari metadata
	tokenStr, err := jwtentity.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Cek token  dari  logout cache
	_, ok := am.cacheService.Get(tokenStr)

	if ok {
		return nil, utils.UnauthenticatedResponse()
	}

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
