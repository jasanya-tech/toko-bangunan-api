package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/helpers/exception"
	"toko-bangunan/models/entities"
)

type UserRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []entities.Supplier
	FindById(ctx context.Context, db *sql.DB, id string) entities.Supplier
	FindByEmail(ctx context.Context, db *sql.DB, email string) entities.User
	Create(ctx context.Context, tx *sql.Tx, user entities.User) entities.User
	Update(ctx context.Context, tx *sql.Tx, supplier entities.Supplier) entities.Supplier
	Delete(ctx context.Context, tx *sql.Tx, id string)
}

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []entities.Supplier {
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

func (repository *UserRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) entities.Supplier {
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

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) entities.User {
	SQL := "SELECT id, username, email , role, image, password, created_at, updated_at FROM users WHERE email = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, email)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	var user entities.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Image, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		return user
	}
	panic(exception.BadRequestError{Message: "invalid email or password"})
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user entities.User) entities.User {
	// email checking
	SQL := "SELECT id, email FROM users WHERE email = ? LIMIT 1"
	rowEmail, errEmail := tx.QueryContext(ctx, SQL, user.Email)
	if errEmail != nil {
		panic(errEmail)
	}
	if rowEmail.Next() {
		panic(exception.BadRequestError{Message: "email is alvailable"})
	}
	defer rowEmail.Close()

	SQL = "INSERT INTO users(id, username, email, role, image, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	_, errInsert := tx.ExecContext(ctx, SQL, user.ID, user.Username, user.Email, user.Role, user.Image, user.Password, user.CreatedAt, user.UpdatedAt)
	if errInsert != nil {
		panic(errInsert)
	}
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, supplier entities.Supplier) entities.Supplier {
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

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {
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
