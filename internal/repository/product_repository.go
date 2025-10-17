package repository

import (
	"database/sql"
)

type IProductRepository interface {
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &productRepository{db: db}
}
