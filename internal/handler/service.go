package handler //*nama package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/abu-umair/ecommerce-go-grpc-be/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

// * membuat methode untuk service handler
func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	panic(errors.New("not implemented")) //?testing error
	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("Hello %s", request.Name),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
