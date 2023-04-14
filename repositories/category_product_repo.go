package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/helpers/exception"
	"toko-bangunan/models/entities"
)

type CategoryProductRepository interface {
	FindALL(ctx context.Context, db *sql.DB) []entities.CategoryProduct
	FindById(ctx context.Context, db *sql.DB, id string) entities.CategoryProduct
	Create(ctx context.Context, tx *sql.Tx, CategoryProduct entities.CategoryProduct) entities.CategoryProduct
	Update(ctx context.Context, tx *sql.Tx, CategoryProduct entities.CategoryProduct) entities.CategoryProduct
	Delete(ctx context.Context, tx *sql.Tx, id string)
}

type CategoryProductRepositoryImpl struct{}

func NewCategoryProductRepository() CategoryProductRepository {
	return &CategoryProductRepositoryImpl{}
}

func (repository *CategoryProductRepositoryImpl) FindALL(ctx context.Context, db *sql.DB) []entities.CategoryProduct {
	SQL := "SELECT id, name, created_at, updated_at FROM product_categories"
	rows, err := db.QueryContext(ctx, SQL)
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
	return categoryProducts
}

func (repository *CategoryProductRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) entities.CategoryProduct {
	SQL := "SELECT id, name, created_at, updated_at FROM product_categories WHERE id=? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	if row.Next() {
		var categoryProduct entities.CategoryProduct
		if err := row.Scan(&categoryProduct.ID, &categoryProduct.Name, &categoryProduct.CreatedAt, &categoryProduct.UpdatedAt); err != nil {
			panic(err)
		}
		return categoryProduct
	}
	panic(exception.NotFoundError{Message: "category product is alvailable"})
}

func (repository *CategoryProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, CategoryProduct entities.CategoryProduct) entities.CategoryProduct {
	SQL := "SELECT id, name FROM product_categories WHERE name=? LIMIT 1"
	row, err := tx.QueryContext(ctx, SQL, CategoryProduct.Name)
	if err != nil {
		panic(err)
	}
	if row.Next() {
		panic(exception.BadRequestError{Message: "name category product is alvailable"})
	}
	SQL = "INSERT INTO product_categories(id, name, created_at, updated_at) VALUES(?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, CategoryProduct.ID, CategoryProduct.Name, CategoryProduct.CreatedAt, CategoryProduct.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return CategoryProduct
}

func (repository *CategoryProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, CategoryProduct entities.CategoryProduct) entities.CategoryProduct {
	SQL := "SELECT id, name FROM product_categories WHERE name = ? AND id <> ? LIMIT 1"
	row, err := tx.QueryContext(ctx, SQL, CategoryProduct.Name, CategoryProduct.ID)
	if err != nil {
		panic(err)
	}
	if row.Next() {
		panic(exception.BadRequestError{Message: "name category product is alvailable"})
	}

	SQL = "UPDATE product_categories SET name=?, updated_at=? WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, CategoryProduct.Name, CategoryProduct.UpdatedAt, CategoryProduct.ID)
	if err != nil {
		panic(err)
	}
	return CategoryProduct
}

func (repository *CategoryProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM product_categories WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	row, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if row == 0 {
		panic(exception.NotFoundError{Message: "category product is alvailable"})
	}
}
