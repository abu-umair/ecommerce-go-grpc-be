package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
)

type ICartRepository interface {
	GetCartByProductAndUserId(ctx context.Context, productId, userId string) (*entity.UserCart, error)
}

type cartRepository struct {
	db *sql.DB
}

func (cr *cartRepository) GetCartByProductAndUserId(ctx context.Context, productId, userId string) (*entity.UserCart, error) {
	row := cr.db.QueryRowContext(
		ctx,
		"SELECT id, product_id, user_id, quantity, created_at, created_by, updated_at, updated_by FROM user_cart WHERE product_id = $1 AND user_id = $2",
		productId,
		userId,
	)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var cartEntity entity.UserCart
	err := row.Scan(
		&cartEntity.Id,
		&cartEntity.ProductId,
		&cartEntity.UserId,
		&cartEntity.Quantity,
		&cartEntity.CreatedAt,
		&cartEntity.CreatedBy,
		&cartEntity.UpdatedAt,
		&cartEntity.UpdatedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return nil, nil
		}

		return nil, err
	}

	return &cartEntity, nil
}

func NewCartRepository(db *sql.DB) ICartRepository {
	return &cartRepository{
		db: db,
	}
}
