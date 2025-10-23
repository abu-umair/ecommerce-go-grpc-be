package grpcmiddleware

import (
	"context"
	"fmt"
	"log"

	jwtentity "github.com/abu-umair/ecommerce-go-grpc-be/internal/entity/jwt"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	gocache "github.com/patrickmn/go-cache"

	"google.golang.org/grpc"
)

type authMiddleware struct {
	cacheService *gocache.Cache
}

func (am *authMiddleware) Middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println(info.FullMethod) //? hanya utk mendapatkan data FullMethod nya saja, 
	log.Println(info.FullMethod)
	if info.FullMethod == "/auth.AuthService/Login" || info.FullMethod == "/auth.AuthService/Register" ||  info.FullMethod == "/product.ProductService/DetailProduct"{
		return handler(ctx, req)
	}

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
	claims, err := jwtentity.GetClaimsFromToken(tokenStr)

	if err != nil {
		return nil, err
	}

	// Sematkan entity ke context
	ctx = claims.SetToContext(ctx)

	res, err := handler(ctx, req)

	return res, err
}

func NewAuthMiddleware(cacheService *gocache.Cache) *authMiddleware {
	return &authMiddleware{
		cacheService: cacheService,
	}
}
