package handler

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/newsletter"
)

type newsletterHandler struct {
	newsletter.UnimplementedNewsletterServiceServer

	newsletterService service.InewsLetterService
}

func (nh *newsletterHandler) SubcribeNewsletter(ctx context.Context, request *newsletter.SubcribeNewsletterRequest) (*newsletter.SubcribeNewsletterResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}
	if validationErrors != nil {
		return &newsletter.SubcribeNewsletterResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}

	res, err := nh.newsletterService.SubcribeNewsletter(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func NewNewsletterHandler(newsletterService service.InewsLetterService) *newsletterHandler {
	return &newsletterHandler{
		newsletterService: newsletterService,
	}
}
