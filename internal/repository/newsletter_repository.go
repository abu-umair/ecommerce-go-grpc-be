package repository

import (
	"context"
	"database/sql"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
)

type InewsLetterRepository interface {
	GetNewsLetterByEmail(ctx context.Context, email string) (*entity.Newsletter, error)
	CreateNewNewsletter(ctx context.Context, newsletter *entity.Newsletter) error
}

type newsLetterRepository struct {
	db *sql.DB
}

func (nr *newsLetterRepository) GetNewsLetterByEmail(ctx context.Context, email string) (*entity.Newsletter, error) {

}
func (nr *newsLetterRepository) CreateNewNewsletter(ctx context.Context, newsletter *entity.Newsletter) error {

}

func NewNewsLetterRespository(db *sql.DB) InewsLetterRepository {
	return &newsLetterRepository{}
