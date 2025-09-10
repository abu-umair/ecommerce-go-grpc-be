package grpcmiddleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func auth_middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// Ambil token  dari metadata
	
	// Cek token  dari  logout cache

	// Parse jwt nya hingga jadi entity

	// Sematkan entity ke context

	res, err := handler(ctx, req)

	return res, err
}
