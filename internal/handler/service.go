package handler //*nama package handler

import (
	"context"
	"fmt"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

// * membuat methode untuk service handler
func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	//? validasi request
	validationErrors, err := utils.CheckValidation(request) //?mirip Validator::make($request->all(), $rules)
	if err != nil {                                         //?mirip if ($validator->fails())
		return nil, err
	}

	if validationErrors != nil {
		return &service.HelloWorldResponse{
			Base: utils.ValidationErrorResponse(validationErrors), //? memindahkan ke sebuah file agar clean dan bisa di reuse
		}, nil
	}

	// panic(errors.New("not implemented")) //?testing error
	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("Hello %s", request.Name), //?mirip return response()->json(["message" => "Hello ".$request->name])
		Base:    utils.SuccessResponse("Success"),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
