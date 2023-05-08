package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/category-product/entities"
)

type CategoryProductRepositoryQuery interface {
	FindALL(ctx context.Context, db *sql.DB, where string) *[]entities.CategoryProduct
	FindById(ctx context.Context, db *sql.DB, id string) (*entities.CategoryProduct, error)
}

type CategoryProductRepositoryQueryImpl struct{}

func (repository *CategoryProductRepositoryQueryImpl) FindALL(ctx context.Context, db *sql.DB, where string) *[]entities.CategoryProduct {
	var rows *sql.Rows
	var err error
	if where != "" {
		SQL := "SELECT id, name, created_at, updated_at FROM product_categories " + where
		// panic(SQL)
		rows, err = db.QueryContext(ctx, SQL)
	} else {
		SQL := "SELECT id, name, created_at, updated_at FROM product_categories"
		rows, err = db.QueryContext(ctx, SQL)
	}
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categoryProducts []entities.CategoryProduct
	for rows.Next() {
		var categoryProduct entities.CategoryProduct

		if err := rows.Scan(&categoryProduct.ID, &categoryProduct.Name, &categoryProduct.CreatedAt, &categoryProduct.UpdatedAt); err != nil {
			panic(err)
		}
		categoryProducts = append(categoryProducts, categoryProduct)
	}
	return &categoryProducts
}

func (repository *CategoryProductRepositoryQueryImpl) FindById(ctx context.Context, db *sql.DB, id string) (*entities.CategoryProduct, error) {
	SQL := "SELECT id, name, created_at, updated_at FROM product_categories WHERE id=? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		var categoryProduct entities.CategoryProduct
		if err := row.Scan(&categoryProduct.ID, &categoryProduct.Name, &categoryProduct.CreatedAt, &categoryProduct.UpdatedAt); err != nil {
			return nil, err
		}
		return &categoryProduct, nil
	}
	return nil, exception.NotFoundError{Message: "category product is alvailable"}
}
