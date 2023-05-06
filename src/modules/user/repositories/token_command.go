package repositories

import (
	"context"
	"database/sql"

	entitytoken "toko-bangunan/src/modules/user/entities"
)

type TokenRepositoryCommand interface {
	TokenCreate(ctx context.Context, tx *sql.Tx, req entitytoken.TokenDetails) (*entitytoken.TokenDetails, error)
	TokenDelete(ctx context.Context, tx *sql.Tx, userId string) error
}

type TokenRepositoryCommandImpl struct{}

func (repo *TokenRepositoryCommandImpl) TokenCreate(ctx context.Context, tx *sql.Tx, req entitytoken.TokenDetails) (*entitytoken.TokenDetails, error) {
	SQL := "INSERT INTO tokens(user_id, refresh_token, expired_token) VALUES(?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, req.UserID, req.Token, req.ExpiresIn)
	if err != nil {
		return nil, err
	}
	return &req, err
}

func (repo *TokenRepositoryCommandImpl) TokenDelete(ctx context.Context, tx *sql.Tx, userId string) error {
	SQL := "DELETE FROM tokens WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	if err != nil {
		return err
	}
	return nil
}
