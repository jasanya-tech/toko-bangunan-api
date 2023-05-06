package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/supplier/entities"
)

type SupplierRepositoryQuery interface {
	FindAll(ctx context.Context, db *sql.DB) *[]entities.Supplier
	FindById(ctx context.Context, db *sql.DB, id string) (*entities.Supplier, error)
}

type SupplierRepositoryQueryImpl struct{}

func NewSupplierRepositoryQueryImpl() SupplierRepositoryQuery {
	return &SupplierRepositoryQueryImpl{}
}

func (repository *SupplierRepositoryQueryImpl) FindAll(ctx context.Context, db *sql.DB) *[]entities.Supplier {
	SQL := "SELECT id, name, email , phone, address, created_at, updated_at FROM suppliers"
	rows, err := db.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	var suppliers []entities.Supplier
	for rows.Next() {
		supplier := entities.Supplier{}
		if err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Email, &supplier.Phone, &supplier.Address, &supplier.CreatedAt, &supplier.UpdatedAt); err != nil {
			panic(err)
		}
		suppliers = append(suppliers, supplier)
	}
	return &suppliers
}

func (repository *SupplierRepositoryQueryImpl) FindById(ctx context.Context, db *sql.DB, id string) (*entities.Supplier, error) {
	SQL := "SELECT id, name, email , phone, address, created_at, updated_at FROM suppliers WHERE id = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var supplier entities.Supplier
	if row.Next() {
		if err := row.Scan(&supplier.ID, &supplier.Name, &supplier.Email, &supplier.Phone, &supplier.Address, &supplier.CreatedAt, &supplier.UpdatedAt); err != nil {
			return nil, err
		}
		return &supplier, nil
	}
	return nil, exception.NotFoundError{Message: "supplier data not available"}
}
