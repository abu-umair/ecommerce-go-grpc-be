package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
)

type InewsLetterRepository interface {
	GetNewsLetterByEmail(ctx context.Context, email string) (*entity.Newsletter, error)
	CreateNewNewsletter(ctx context.Context, newsletter *entity.Newsletter) error
}

type newsLetterRepository struct {
	db *sql.DB
}

//? menambahkan querynya 'GetNewsLetterByEmail'
func (nr *newsLetterRepository) GetNewsLetterByEmail(ctx context.Context, email string) (*entity.Newsletter, error) {
	row := nr.db.QueryRowContext(
		ctx,
		"SELECT id FROM newsletter WHERE email = $1 AND is_deleted = false",
		email, //?email sebagai argumen (bkn ID)
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var newsletter entity.Newsletter
	err := row.Scan(
		&newsletter.Id,
	)
	if err != nil {//?jika ada error atau tidak ketemu
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &newsletter, nil //?jika datanya ketemu
}

func (nr *newsLetterRepository) CreateNewNewsletter(ctx context.Context, newsletter *entity.Newsletter) error {

}

func NewNewsLetterRespository(db *sql.DB) InewsLetterRepository {
	return &newsLetterRepository{
		db: db,
	}
}
