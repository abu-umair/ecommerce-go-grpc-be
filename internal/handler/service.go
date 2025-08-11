package handler //*nama package handler

import "github.com/abu-umair/ecommerce-go-grpc-be/pb/service"

type IServiceHandler interface{}

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func NewServiceHandler() IServiceHandler {
	return &serviceHandler{}
}
