package handler //*nama package handler

import (
	"context"
	"fmt"

	"buf.build/go/protovalidate"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

// * membuat methode untuk service handler
func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	if err := protovalidate.Validate(request); err != nil { //?jika ada error maka akan mereturn error
		return nil, status.Errorf(codes.InvalidArgument, "Validation error %v", err)
	}

	// panic(errors.New("not implemented")) //?testing error
	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("Hello %s", request.Name),
		Base:    utils.SuccessResponse("Success"),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
