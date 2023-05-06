package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	userentity "toko-bangunan/src/modules/user/entities"
)

type UserRepositoryCommand interface {
	Create(ctx context.Context, tx *sql.Tx, user *userentity.User) (*userentity.User, error)
}

type UserRepositoryCommandImpl struct{}

func (repo *UserRepositoryCommandImpl) Create(ctx context.Context, tx *sql.Tx, user *userentity.User) (*userentity.User, error) {
	SQL := "SELECT id, email FROM users WHERE email = ?"
	row, err := tx.QueryContext(ctx, SQL, user.Email)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		return nil, exception.BadRequestError{Message: "e-mail already registered"}
	}
	defer row.Close()
	SQL = "INSERT INTO users(id, username, email, role, image, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	_, errInsert := tx.ExecContext(ctx, SQL, user.ID, user.Username, user.Email, user.Role, user.Image, user.Password, user.CreatedAt, user.UpdatedAt)
	if errInsert != nil {
		return nil, errInsert
	}
	return user, nil
}
