package repositories

import (
	"context"
	"database/sql"

	"toko-bangunan/internal/protocols/http/exception"
	entitytoken "toko-bangunan/src/modules/user/entities"
)

type TokenRepositoryQuery interface {
	FindTokenByUserId(ctx context.Context, db *sql.DB, userId string) (*entitytoken.TokenDetails, error)
}

type TokenRepositoryQueryImpl struct{}

func (repo *TokenRepositoryQueryImpl) FindTokenByUserId(ctx context.Context, db *sql.DB, userId string) (*entitytoken.TokenDetails, error) {
	SQL := "SELECT user_id, refresh_token, expired_token FROM tokens WHERE user_id = ? LIMIT 1"
	row, err := db.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var token entitytoken.TokenDetails
	if row.Next() {
		if err := row.Scan(&token.UserID, &token.Token, &token.ExpiresIn); err != nil {
			return nil, err
		}
		return &token, nil
	}
	return nil, exception.Forbidden{Message: "invalid token"}
}
