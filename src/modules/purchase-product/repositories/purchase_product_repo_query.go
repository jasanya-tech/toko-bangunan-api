package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/purchase-product/entities"
)

type PurchaseProductRepositoryQuery interface {
	FindAll(ctx context.Context, db *sql.DB) *[]entities.PurchaseProduct
	FindById(ctx context.Context, db *sql.DB, id string) (*entities.PurchaseProduct, error)
}

type PurchaseProductRepositoryQueryImpl struct{}

func (repository *PurchaseProductRepositoryQueryImpl) FindAll(ctx context.Context, db *sql.DB) *[]entities.PurchaseProduct {
	SQL := `SELECT id, id_product, purchase_amount, purchase_price, total_purchase, status, created_at, updated_at FROM purchase_products`
	rows, err := db.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var purchaseProducts []entities.PurchaseProduct
	for rows.Next() {
		purchaseProduct := entities.PurchaseProduct{}
		if err := rows.Scan(&purchaseProduct.ID, &purchaseProduct.IdProduct, &purchaseProduct.PurchaseAmount, &purchaseProduct.PurchasePrice, &purchaseProduct.PurchaseTotal, &purchaseProduct.Status, &purchaseProduct.CreatedAt, &purchaseProduct.UpdatedAt); err != nil {
			panic(err)
		}
		purchaseProducts = append(purchaseProducts, purchaseProduct)
	}
	return &purchaseProducts
}

func (repository *PurchaseProductRepositoryQueryImpl) FindById(ctx context.Context, db *sql.DB, id string) (*entities.PurchaseProduct, error) {
	SQL := `SELECT id, id_product, purchase_amount, purchase_price, total_purchase, status, created_at, updated_at FROM purchase_products WHERE id = ? LIMIT 1`
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var purchaseProduct entities.PurchaseProduct
	if row.Next() {
		if err := row.Scan(&purchaseProduct.ID, &purchaseProduct.IdProduct, &purchaseProduct.PurchaseAmount, &purchaseProduct.PurchasePrice, &purchaseProduct.PurchaseTotal, &purchaseProduct.Status, &purchaseProduct.CreatedAt, &purchaseProduct.UpdatedAt); err != nil {
			return nil, err
		}
		return &purchaseProduct, nil
	}
	return nil, exception.NotFoundError{Message: "purchase product data not available"}
}
