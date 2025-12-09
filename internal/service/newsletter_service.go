package service

import (
	"context"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/newsletter"
	"github.com/google/uuid"
)

type InewsLetterService interface {
	SubcribeNewsletter(ctx context.Context, request *newsletter.SubcribeNewsletterRequest) (*newsletter.SubcribeNewsletterResponse, error)
}

type newsletterService struct {
}

func (ns *newsletterService) SubcribeNewsletter(ctx context.Context, request *newsletter.SubcribeNewsletterRequest) (*newsletter.SubcribeNewsletterResponse, error) {
	//* cek database, apakah email sudah ada/terdaftar atau belum (jika Sudah terdaftar maka return success dan tidak disimpan di DB)
	

	//* insert ke db (jika email belum terdaftar)
	
}

func NewNewsLetterService(newsletterRepository repository.InewsLetterRepository) InewsLetterService {
	return &newsletterService{
		newsletterRepository: newsletterRepository,
	}
}
