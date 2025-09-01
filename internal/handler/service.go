package handler //*nama package handler

import (
	"context"
	"errors"
	"fmt"

	"buf.build/go/protovalidate"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/common"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/service"
)

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

// * membuat methode untuk service handler
func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	 //? validasi request
	if err := protovalidate.Validate(request); //?validator.Validate(request) cek request sesuai aturan di .proto.
	err != nil {                               //?jika ada error maka akan mereturn error Analogi Laravel → mirip kalau kita bikin custom validator class lalu panggil Validator::make($data, $rules).
		var validationError *protovalidate.ValidationError //?Kalau gagal → return *protovalidate.ValidationError.

		if errors.As(err, &validationError) {
			var validationErrorResponse []*common.ValidationError = make([]*common.ValidationError, 0)
			for _, violation := range validationError.Violations { //? Di Laravel, kalau validasi gagal, kita bisa ambil $errors->first('field') atau $errors->all()., Di Go, validationError.Violations itu array yang mirip Laravel: $errors->all() atau $errors->toArray()
				//? Di Laravel, framework otomatis ubah array error ke JSON.
				//? Di Go, kita harus manual mapping violation → response struct.
				validationErrorResponse = append(validationErrorResponse, &common.ValidationError{
					Field:   *violation.Proto.Field.Elements[0].FieldName, //?FieldPath → nama field (name, email)
					Message: *violation.Proto.Message,                     //?Message → mirip pesan error di Laravel (The name field is required.).
				})
			}
			return &service.HelloWorldResponse{
				Base: &common.BaseResponse{
					ValidationError: validationErrorResponse,
					StatusCode:      400,
					Message:         "Validation Error",
					IsError:         true,
				},
			}, nil
		}
		return nil, err

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
