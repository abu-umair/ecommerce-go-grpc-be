package service

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/newsletter"
)

type InewsLetterService interface {
	SubcribeNewsletter(ctx context.Context, request *newsletter.SubcribeNewsletterRequest) (*newsletter.SubcribeNewsletterResponse, error)
}

type newsletterService struct {
	newsletterRepository repository.InewsLetterRepository
}

func (ns *newsletterService) SubcribeNewsletter(ctx context.Context, request *newsletter.SubcribeNewsletterRequest) (*newsletter.SubcribeNewsletterResponse, error) {
	//* cek database, apakah email sudah ada/terdaftar atau belum (jika Sudah terdaftar maka return success dan tidak disimpan di DB)

	//* insert ke db (jika email belum terdaftar)

}

func NewNewsLetterService(newsletterRepository repository.InewsLetterRepository) InewsLetterService { //?newsletterRepository (dari `type newsletterService struct`)
	return &newsletterService{
		newsletterRepository: newsletterRepository, //?diisi dari repository (dari `type newsletterService struct`)
	}
}
