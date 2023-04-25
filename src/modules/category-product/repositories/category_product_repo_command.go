package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/category-product/entities"
)

type CategoryProductRepositoryCommand interface {
	Create(ctx context.Context, tx *sql.Tx, CategoryProduct *entities.CategoryProduct) (*entities.CategoryProduct, error)
	Update(ctx context.Context, tx *sql.Tx, CategoryProduct *entities.CategoryProduct) (*entities.CategoryProduct, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}

type CategoryProductRepositoryCommandImpl struct{}

func (repository *CategoryProductRepositoryCommandImpl) Create(ctx context.Context, tx *sql.Tx, categoryProduct *entities.CategoryProduct) (*entities.CategoryProduct, error) {
	SQL := "SELECT id, name FROM product_categories WHERE name=? LIMIT 1"
	row, err := tx.QueryContext(ctx, SQL, categoryProduct.Name)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		return nil, exception.BadRequestError{Message: "name category product is alvailable"}
	}
	SQL = "INSERT INTO product_categories(id, name, created_at, updated_at) VALUES(?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, categoryProduct.ID, categoryProduct.Name, categoryProduct.CreatedAt, categoryProduct.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return categoryProduct, nil
}

func (repository *CategoryProductRepositoryCommandImpl) Update(ctx context.Context, tx *sql.Tx, categoryProduct *entities.CategoryProduct) (*entities.CategoryProduct, error) {
	SQL := "SELECT id, name FROM product_categories WHERE name = ? AND id <> ? LIMIT 1"
	row, err := tx.QueryContext(ctx, SQL, categoryProduct.Name, categoryProduct.ID)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		return nil, exception.BadRequestError{Message: "name category product is alvailable"}
	}

	SQL = "UPDATE product_categories SET name=?, updated_at=? WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, categoryProduct.Name, categoryProduct.UpdatedAt, categoryProduct.ID)
	if err != nil {
		return nil, err
	}
	return categoryProduct, nil
}

func (repository *CategoryProductRepositoryCommandImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM product_categories WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return exception.NotFoundError{Message: "category product is alvailable"}
	}
	return nil
}
