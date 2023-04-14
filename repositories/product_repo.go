package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/helpers/exception"
	"toko-bangunan/models/entities"
)

type ProductRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []entities.Product
	FindById(ctx context.Context, db *sql.DB, id string) entities.Product
	Create(ctx context.Context, tx *sql.Tx, supplier entities.Product) entities.Product
	Update(ctx context.Context, tx *sql.Tx, supplier entities.Product) entities.Product
	Delete(ctx context.Context, tx *sql.Tx, id string)
}

type ProductRepositoryImpl struct{}

func NewProductRepositoryImpl() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []entities.Product {
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
	return products
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) entities.Product {
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
		panic(err)
	}
	defer row.Close()

	var product entities.Product
	if row.Next() {
		if err := row.Scan(
			&product.ID, &product.Name, &product.SellingPrice, &product.PurchasePrice, &product.StockProduct, &product.Image, &product.CreatedAt, &product.UpdatedAt,
			&product.Supplier.ID, &product.Supplier.Name, &product.Supplier.Email, &product.Supplier.Phone, &product.Supplier.Address, &product.Supplier.CreatedAt, &product.Supplier.UpdatedAt,
			&product.ProductCategory.ID, &product.ProductCategory.Name, &product.ProductCategory.CreatedAt, &product.ProductCategory.UpdatedAt,
		); err != nil {
			panic(err)
		}
		return product
	}
	panic(exception.NotFoundError{Message: "product data not available"})
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product entities.Product) entities.Product {
	// name checking
	SQL := "SELECT id, name FROM products WHERE name = ? LIMIT 1"
	rowName, errName := tx.QueryContext(ctx, SQL, product.Name)
	if errName != nil {
		panic(errName)
	}
	if rowName.Next() {
		panic(exception.BadRequestError{Message: "name product is alvailable"})
	}
	defer rowName.Close()

	SQL = "INSERT INTO products(id, supplier_id, product_category_id, name, selling_price, purchase_price, stock_product, image, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errInsert := tx.ExecContext(ctx, SQL, product.ID, product.Supplier.ID, product.ProductCategory.ID, product.Name, product.SellingPrice, product.PurchasePrice, product.StockProduct, product.Image, product.CreatedAt, product.UpdatedAt)
	if errInsert != nil {
		panic(errInsert)
	}
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entities.Product) entities.Product {
	// name checking
	SQL := "SELECT id, name FROM products WHERE name = ? AND id <> ? LIMIT 1"
	rowName, errName := tx.QueryContext(ctx, SQL, product.Name, product.ID)
	if errName != nil {
		panic(errName)
	}
	if rowName.Next() {
		panic(exception.BadRequestError{Message: "name product is alvailable"})
	}
	defer rowName.Close()

	SQL = "UPDATE products SET supplier_id = ?, product_category_id = ?, name = ?, selling_price = ?, purchase_price = ?, stock_product = ?, image = ?, created_at = ?, updated_at = ? WHERE id = ?"
	_, errInsert := tx.ExecContext(ctx, SQL, product.Supplier.ID, product.ProductCategory.ID, product.Name, product.SellingPrice, product.PurchasePrice, product.StockProduct, product.Image, product.CreatedAt, product.UpdatedAt, product.ID)
	if errInsert != nil {
		panic(errInsert)
	}
	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM products WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}
