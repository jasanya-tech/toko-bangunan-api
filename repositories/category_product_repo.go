package repositories

import (
	"context"
	"database/sql"

	"github.com/SyaibanAhmadRamadhan/toko-bangunan/models/entities"
)

type CategoryProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, CategoryProduct entities.CategoryProduct) entities.CategoryProduct
}

type CategoryProductRepositoryImpl struct{}

func NewCategoryProductRepository() CategoryProductRepository {
	return &CategoryProductRepositoryImpl{}
}

func (repository *CategoryProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, CategoryProduct entities.CategoryProduct) entities.CategoryProduct {
	SQL := "INSERT INTO product_categories(id, name, created_at, updated_at) VALUES(?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, CategoryProduct.ID, CategoryProduct.Name, CategoryProduct.CreatedAt, CategoryProduct.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return CategoryProduct
}
