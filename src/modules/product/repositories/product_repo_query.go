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

func NewProductRepositoryQueryImpl() ProductRepositoryQuery {
	return &ProductRepositoryQueryImpl{}
}

func (repository *ProductRepositoryQueryImpl) FindAll(ctx context.Context, db *sql.DB) *[]entities.Product {
	SQL := `
			SELECT products.id, products.name, products.selling_price , products.purchase_price, products.stock_product, products.image, products.created_at, products.updated_at, products.supplier_id,
			suppliers.name, suppliers.email, suppliers.phone, suppliers.address, suppliers.created_at, suppliers.updated_at,
			products.product_category_id, product_categories.name, product_categories.created_at, product_categories.updated_at
			FROM products
			INNER JOIN suppliers ON products.supplier_id = suppliers.id
			INNER JOIN product_categories ON products.product_category_id = product_categories.id
			`
	rows, err := db.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		product := entities.Product{}
		if err := rows.Scan(
			&product.ID, &product.Name, &product.SellingPrice, &product.PurchasePrice, &product.StockProduct, &product.Image, &product.CreatedAt, &product.UpdatedAt,
			&product.Supplier.ID, &product.Supplier.Name, &product.Supplier.Email, &product.Supplier.Phone, &product.Supplier.Address, &product.Supplier.CreatedAt, &product.Supplier.UpdatedAt,
			&product.ProductCategory.ID, &product.ProductCategory.Name, &product.ProductCategory.CreatedAt, &product.ProductCategory.UpdatedAt,
		); err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return &products
}

func (repository *ProductRepositoryQueryImpl) FindById(ctx context.Context, db *sql.DB, id string) (*entities.Product, error) {
	SQL := `
			SELECT products.id, products.name, products.selling_price , products.purchase_price, products.stock_product, products.image, products.created_at, products.updated_at, products.supplier_id,
			suppliers.name, suppliers.email, suppliers.phone, suppliers.address, suppliers.created_at, suppliers.updated_at,
			products.product_category_id, product_categories.name, product_categories.created_at, product_categories.updated_at
			FROM products
			INNER JOIN suppliers ON products.supplier_id = suppliers.id
			INNER JOIN product_categories ON products.product_category_id = product_categories.id
			WHERE products.id = ? LIMIT 1
			`
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var product entities.Product
	if row.Next() {
		if err := row.Scan(
			&product.ID, &product.Name, &product.SellingPrice, &product.PurchasePrice, &product.StockProduct, &product.Image, &product.CreatedAt, &product.UpdatedAt,
			&product.Supplier.ID, &product.Supplier.Name, &product.Supplier.Email, &product.Supplier.Phone, &product.Supplier.Address, &product.Supplier.CreatedAt, &product.Supplier.UpdatedAt,
			&product.ProductCategory.ID, &product.ProductCategory.Name, &product.ProductCategory.CreatedAt, &product.ProductCategory.UpdatedAt,
		); err != nil {
			return nil, err
		}
		return &product, nil
	}
	return nil, exception.NotFoundError{Message: "product data not available"}
}
