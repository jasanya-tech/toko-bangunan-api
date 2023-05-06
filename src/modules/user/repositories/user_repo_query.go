package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	userentity "toko-bangunan/src/modules/user/entities"
)

type UserRepositoryQuery interface {
	FindById(ctx context.Context, db *sql.DB, id string) (*userentity.User, error)
	FindByEmail(ctx context.Context, db *sql.DB, email string) (*userentity.User, error)
}

type UserRepositoryQueryImpl struct{}

func (repo *UserRepositoryQueryImpl) FindById(ctx context.Context, db *sql.DB, id string) (*userentity.User, error) {
	SQL := "SELECT id, username, email , role, image, password, created_at, updated_at FROM users WHERE id = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user userentity.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Image, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		return &user, nil
	} else {
		return nil, exception.NotFoundError{Message: "user not found"}
	}
}

func (repo *UserRepositoryQueryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) (*userentity.User, error) {
	SQL := "SELECT id, username, email , role, image, password, created_at, updated_at FROM users WHERE email = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, email)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user userentity.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Image, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		return &user, nil
	} else {
		return nil, exception.NotFoundError{Message: "user not found"}
	}
}
