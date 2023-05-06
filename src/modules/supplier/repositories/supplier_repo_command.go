package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/src/modules/supplier/entities"
)

type SupplierRepositoryCommand interface {
	Create(ctx context.Context, tx *sql.Tx, supplier *entities.Supplier) (*entities.Supplier, error)
	Update(ctx context.Context, tx *sql.Tx, supplier *entities.Supplier) (*entities.Supplier, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}

type SupplierRepositoryCommandImpl struct{}

func NewSupplierRepositoryCommandImpl() SupplierRepositoryCommand {
	return &SupplierRepositoryCommandImpl{}
}

func (repository *SupplierRepositoryCommandImpl) Create(ctx context.Context, tx *sql.Tx, supplier *entities.Supplier) (*entities.Supplier, error) {
	// email checking
	SQL := "SELECT id, email FROM suppliers WHERE email = ? LIMIT 1"
	rowEmail, errEmail := tx.QueryContext(ctx, SQL, supplier.Email)
	if errEmail != nil {
		return nil, errEmail
	}
	if rowEmail.Next() {
		return nil, exception.BadRequestError{Message: "email is alvailable"}
	}
	defer rowEmail.Close()

	// phone checking
	SQL = "SELECT id, phone FROM suppliers WHERE phone = ? LIMIT 1"
	rowPhone, errPhone := tx.QueryContext(ctx, SQL, supplier.Phone)
	if errPhone != nil {
		return nil, errPhone
	}
	if rowPhone.Next() {
		return nil, exception.BadRequestError{Message: "phone is alvailable"}
	}
	defer rowPhone.Close()

	SQL = "INSERT INTO suppliers(id, name, email, phone, address, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)"
	_, errInsert := tx.ExecContext(ctx, SQL, supplier.ID, supplier.Name, supplier.Email, supplier.Phone, supplier.Address, supplier.CreatedAt, supplier.UpdatedAt)
	if errInsert != nil {
		return nil, errInsert
	}
	return supplier, nil
}

func (repository *SupplierRepositoryCommandImpl) Update(ctx context.Context, tx *sql.Tx, supplier *entities.Supplier) (*entities.Supplier, error) {
	// email checking
	SQL := "SELECT id, email FROM suppliers WHERE email = ? AND id <> ?"
	rowEmail, errEmail := tx.QueryContext(ctx, SQL, supplier.Email, supplier.ID)
	if errEmail != nil {
		return nil, errEmail
	}
	if rowEmail.Next() {
		return nil, exception.BadRequestError{Message: "email is alvailable"}
	}
	defer rowEmail.Close()

	// phone checking
	SQL = "SELECT id, phone FROM suppliers WHERE phone = ? AND id <> ? LIMIT 1"
	rowPhone, errPhone := tx.QueryContext(ctx, SQL, supplier.Phone, supplier.ID)
	if errPhone != nil {
		return nil, errPhone
	}
	if rowPhone.Next() {
		return nil, exception.BadRequestError{Message: "phone is alvailable"}
	}
	defer rowPhone.Close()

	SQL = "UPDATE suppliers SET name = ?, email = ?, phone = ?, address = ?, updated_at = ? WHERE id = ?"
	_, errUpdate := tx.ExecContext(ctx, SQL, supplier.Name, supplier.Email, supplier.Phone, supplier.Address, supplier.UpdatedAt, supplier.ID)
	if errUpdate != nil {
		return nil, errUpdate
	}
	return supplier, nil
}

func (repository *SupplierRepositoryCommandImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM suppliers WHERE id = ?"
	row, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	res, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return exception.NotFoundError{Message: "supplier is alvailable"}
	}
	SQL = "DELETE FROM products WHERE supplier_id = ?"
	if _, err := tx.ExecContext(ctx, SQL, id); err != nil {
		return err
	}
	return nil
}
