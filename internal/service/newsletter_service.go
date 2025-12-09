package service

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/pb/newsletter"
)

type InewsLetterService interface {
	SubcribeNewsletter(ctx context.Context, request *newsletter.SubcribeNewsletterRequest) (*newsletter.SubcribeNewsletterResponse, error)
}
