package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/product/entities"
)

type ProductRepositoryQuery interface {
	FindAll(ctx context.Context, db *sql.DB) *[]entities.Product
	FindById(ctx context.Context, db *sql.DB, id string) (*entities.Product, error)
}

type ProductRepositoryQueryImpl struct{}

func (repository *ProductRepositoryQueryImpl) FindAll(ctx context.Context, db *sql.DB) *[]entities.Product {
	SQL := `SELECT id, supplier_id, product_category_id, name, selling_price, stock_product, image, created_at, updated_at FROM products`
	rows, err := db.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		product := entities.Product{}
		if err := rows.Scan(
			&product.ID, &product.SupplierId, &product.ProductCategoryId, &product.Name, &product.SellingPrice, &product.StockProduct, &product.Image, &product.CreatedAt, &product.UpdatedAt,
		); err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return &products
}

func (repository *ProductRepositoryQueryImpl) FindById(ctx context.Context, db *sql.DB, id string) (*entities.Product, error) {
	SQL := `SELECT id, supplier_id, product_category_id, name, selling_price, stock_product, image, created_at, updated_at FROM products WHERE id = ? LIMIT 1`
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var product entities.Product

	if row.Next() {
		if err := row.Scan(
			&product.ID, &product.SupplierId, &product.ProductCategoryId, &product.Name, &product.SellingPrice, &product.StockProduct, &product.Image, &product.CreatedAt, &product.UpdatedAt,
		); err != nil {
			panic(err)
		}
		return &product, nil
	}
	return nil, exception.NotFoundError{Message: "product data not available"}
}
