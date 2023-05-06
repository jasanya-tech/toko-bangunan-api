package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/product/entities"
)

type ProductRepositoryCommand interface {
	Create(ctx context.Context, tx *sql.Tx, product *entities.Product) (*entities.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product *entities.Product) (*entities.Product, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}

type ProductRepositoryCommandImpl struct{}

func (repository *ProductRepositoryCommandImpl) Create(ctx context.Context, tx *sql.Tx, product *entities.Product) (*entities.Product, error) {
	// name checking
	SQL := "SELECT id, name FROM products WHERE name = ? LIMIT 1"
	rowName, errName := tx.QueryContext(ctx, SQL, product.Name)
	if errName != nil {
		return nil, errName
	}
	if rowName.Next() {
		return nil, exception.BadRequestError{Message: "name product is alvailable"}
	}
	defer rowName.Close()

	SQL = "INSERT INTO products(id, supplier_id, product_category_id, name, selling_price, image, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	_, errInsert := tx.ExecContext(ctx, SQL, product.ID, product.SupplierId, product.ProductCategoryId, product.Name, product.SellingPrice, product.Image, product.CreatedAt, product.UpdatedAt)
	if errInsert != nil {
		return nil, errInsert
	}
	return product, nil
}

func (repository *ProductRepositoryCommandImpl) Update(ctx context.Context, tx *sql.Tx, product *entities.Product) (*entities.Product, error) {
	// name checking
	SQL := "SELECT id, name FROM products WHERE name = ? AND id <> ? LIMIT 1"
	rowName, errName := tx.QueryContext(ctx, SQL, product.Name, product.ID)
	if errName != nil {
		return nil, errName
	}
	if rowName.Next() {
		return nil, exception.BadRequestError{Message: "name product is alvailable"}
	}
	defer rowName.Close()

	SQL = "UPDATE products SET supplier_id = ?, product_category_id = ?, name = ?, selling_price = ?, image = ?, created_at = ?, updated_at = ? WHERE id = ?"
	_, errInsert := tx.ExecContext(ctx, SQL, product.SupplierId, product.ProductCategoryId, product.Name, product.SellingPrice, product.Image, product.CreatedAt, product.UpdatedAt, product.ID)
	if errInsert != nil {
		return nil, errInsert
	}
	return product, nil
}

func (repository *ProductRepositoryCommandImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM products WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return exception.NotFoundError{Message: "product purchase not alvailable"}
	}
	return nil
}
