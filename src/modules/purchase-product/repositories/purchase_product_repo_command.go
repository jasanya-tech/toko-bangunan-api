package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/purchase-product/entities"
)

type PurchaseProductRepositoryCommand interface {
	Create(ctx context.Context, tx *sql.Tx, purchase *entities.PurchaseProduct) (*entities.PurchaseProduct, error)
	Update(ctx context.Context, tx *sql.Tx, purchase *entities.PurchaseProduct) (*entities.PurchaseProduct, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}

type PurchaseProductRepositoryCommandImpl struct{}

func (repository *PurchaseProductRepositoryCommandImpl) Create(ctx context.Context, tx *sql.Tx, purchase *entities.PurchaseProduct) (*entities.PurchaseProduct, error) {
	SQL := "INSERT INTO purchase_products(id, id_product, purchase_amount, purchase_price, total_purchase, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	_, errInsert := tx.ExecContext(ctx, SQL, purchase.ID, purchase.IdProduct, purchase.PurchaseAmount, purchase.PurchasePrice, purchase.PurchaseTotal, purchase.Status, purchase.CreatedAt, purchase.UpdatedAt)
	if errInsert != nil {
		return nil, errInsert
	}

	return purchase, nil
}

func (repository *PurchaseProductRepositoryCommandImpl) Update(ctx context.Context, tx *sql.Tx, purchase *entities.PurchaseProduct) (*entities.PurchaseProduct, error) {
	SQL := "UPDATE purchase_products SET id_product = ?, purchase_amount = ?, purchase_price = ?, total_purchase = ?, status = ?, created_at = ?, updated_at = ? WHERE id = ?"
	_, errInsert := tx.ExecContext(ctx, SQL, purchase.IdProduct, purchase.PurchaseAmount, purchase.PurchasePrice, purchase.PurchaseTotal, purchase.Status, purchase.CreatedAt, purchase.UpdatedAt, purchase.ID)
	if errInsert != nil {
		return nil, errInsert
	}
	return purchase, nil
}

func (repository *PurchaseProductRepositoryCommandImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM purchase_products WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return exception.NotFoundError{Message: "purchase product not alvailable"}
	}
	return nil
}
