package utils //?agar bisa direuse

import "github.com/abu-umair/ecommerce-go-grpc-be/pb/common"

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message,
	}
}
