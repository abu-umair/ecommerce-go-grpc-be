package utils

import (
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/common"
	"google.golang.org/protobuf/proto"
)

func CheckValidation(req proto.Message) ([]*common.ValidationError, error) {
	
}
