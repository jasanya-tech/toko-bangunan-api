package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/helpers/exception"
	"toko-bangunan/models/entities"
)

type SupplierRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []entities.Supplier
	FindById(ctx context.Context, db *sql.DB, id string) entities.Supplier
	Create(ctx context.Context, tx *sql.Tx, supplier entities.Supplier) entities.Supplier
	Update(ctx context.Context, tx *sql.Tx, supplier entities.Supplier) entities.Supplier
	Delete(ctx context.Context, tx *sql.Tx, id string)
}

type SupplierRepositoryImpl struct{}

func NewSupplierRepositoryImpl() SupplierRepository {
	return &SupplierRepositoryImpl{}
}

func (repository *SupplierRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []entities.Supplier {
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
	return suppliers
}

func (repository *SupplierRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) entities.Supplier {
	SQL := "SELECT id, name, email , phone, address, created_at, updated_at FROM suppliers WHERE id = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	var supplier entities.Supplier
	if row.Next() {
		if err := row.Scan(&supplier.ID, &supplier.Name, &supplier.Email, &supplier.Phone, &supplier.Address, &supplier.CreatedAt, &supplier.UpdatedAt); err != nil {
			panic(err)
		}
		return supplier
	}
	panic(exception.NotFoundError{Message: "supplier data not available"})
}

func (repository *SupplierRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, supplier entities.Supplier) entities.Supplier {
	// email checking
	SQL := "SELECT id, email FROM suppliers WHERE email = ? LIMIT 1"
	rowEmail, errEmail := tx.QueryContext(ctx, SQL, supplier.Email)
	if errEmail != nil {
		panic(errEmail)
	}
	if rowEmail.Next() {
		panic(exception.BadRequestError{Message: "email is alvailable"})
	}
	defer rowEmail.Close()

	// phone checking
	SQL = "SELECT id, phone FROM suppliers WHERE phone = ? LIMIT 1"
	rowPhone, errPhone := tx.QueryContext(ctx, SQL, supplier.Phone)
	if errPhone != nil {
		panic(errPhone)
	}
	if rowPhone.Next() {
		panic(exception.BadRequestError{Message: "phone is alvailable"})
	}
	defer rowPhone.Close()

	SQL = "INSERT INTO suppliers(id, name, email, phone, address, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)"
	_, errInsert := tx.ExecContext(ctx, SQL, supplier.ID, supplier.Name, supplier.Email, supplier.Phone, supplier.Address, supplier.CreatedAt, supplier.UpdatedAt)
	if errInsert != nil {
		panic(errInsert)
	}
	return supplier
}

func (repository *SupplierRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, supplier entities.Supplier) entities.Supplier {
	// email checking
	SQL := "SELECT id, email FROM suppliers WHERE email = ? AND id <> ?"
	rowEmail, errEmail := tx.QueryContext(ctx, SQL, supplier.Email, supplier.ID)
	if errEmail != nil {
		panic(errEmail)
	}
	if rowEmail.Next() {
		panic(exception.BadRequestError{Message: "email is alvailable"})
	}
	defer rowEmail.Close()

	// phone checking
	SQL = "SELECT id, phone FROM suppliers WHERE phone = ? AND id <> ? LIMIT 1"
	rowPhone, errPhone := tx.QueryContext(ctx, SQL, supplier.Phone, supplier.ID)
	if errPhone != nil {
		panic(errPhone)
	}
	if rowPhone.Next() {
		panic(exception.BadRequestError{Message: "phone is alvailable"})
	}
	defer rowPhone.Close()

	SQL = "UPDATE suppliers SET name = ?, email = ?, phone = ?, address = ?, updated_at = ? WHERE id = ?"
	_, errUpdate := tx.ExecContext(ctx, SQL, supplier.Name, supplier.Email, supplier.Phone, supplier.Address, supplier.UpdatedAt, supplier.ID)
	if errUpdate != nil {
		panic(errUpdate)
	}
	return supplier
}

func (repository *SupplierRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM suppliers WHERE id = ?"
	row, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	res, err := row.RowsAffected()
	if err != nil {
		panic(err)
	}
	if res == 0 {
		panic(exception.NotFoundError{Message: "supplier is alvailable"})
	}
	SQL = "DELETE FROM products WHERE supplier_id = ?"
	if _, err := tx.ExecContext(ctx, SQL, id); err != nil {
		panic(err)
	}
}
