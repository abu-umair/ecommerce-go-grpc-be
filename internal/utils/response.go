package utils //?agar bisa direuse

import "github.com/abu-umair/ecommerce-go-grpc-be/pb/common"

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message,
	}
}

func ValidationErrorResponse(validationErrors []*common.ValidationError) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode:      400,
		Message:         "Validation error",
		IsError:         true,
		ValidationError: validationErrors, //?mirip return JSON $validator->errors()->toArray()
	}
}
