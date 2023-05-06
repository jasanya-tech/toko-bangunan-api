package auth

import (
	"context"
	"database/sql"
	"encoding/base64"
	"time"

	"toko-bangunan/config"
	"toko-bangunan/infrastructures/db/transaction"
	tokendto "toko-bangunan/internal/utils/auth/dto"
	rsaheader "toko-bangunan/internal/utils/rsa"
	"toko-bangunan/src/modules/user/entities"
	userrepo "toko-bangunan/src/modules/user/repositories"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

type JwtToken interface {
	SignRsa(claims jwt.MapClaims) tokendto.Token
}

type JwtTokenImpl struct {
	jwtTokenTimeExp        time.Duration
	jwtTokenRefreshTimeExp time.Duration
	UserRepo               userrepo.Repositories
	DB                     *sql.DB
}

func NewJwtTokenImpl(userRepo userrepo.Repositories, db *sql.DB) *JwtTokenImpl {
	jwtTokenDuration := time.Duration(config.Get().Auth.JwtToken.AccessToken.Expired)
	jwtTokenRefreshDuration := time.Duration(config.Get().Auth.JwtToken.RefreshToken.Expired)
	return &JwtTokenImpl{
		jwtTokenTimeExp:        jwtTokenDuration,
		jwtTokenRefreshTimeExp: jwtTokenRefreshDuration,
		UserRepo:               userRepo,
		DB:                     db,
	}
}

func (o JwtTokenImpl) SignRsa(claims jwt.MapClaims) tokendto.Token {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(o.jwtTokenTimeExp * time.Minute).Unix()
	if claims["id"] == nil {
		return tokendto.Token{}
	}

	token := jwt.New(jwt.SigningMethodRS256)
	// setup user
	_, checkExp := claims["exp"]
	_, checkIat := claims["exp"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}
	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}
	claims["token_type"] = "access_token"

	token.Claims = claims
	authToken := new(tokendto.Token)

	decodePrivateKeyAccess, _ := base64.StdEncoding.DecodeString(config.Get().Auth.JwtToken.AccessToken.PrivateKey)
	privateRsaAccess, err := rsaheader.ReadPrivateKeyFromEnv(decodePrivateKeyAccess)
	if err != nil {
		log.Err(err).Msg("cannot read access private key token rsa from env")
		return tokendto.Token{}
	}

	tokenString, err := token.SignedString(privateRsaAccess)
	if err != nil {
		log.Err(err).Msg("err read access private key rsa")
		return tokendto.Token{}
	}

	authToken.Token = tokenString
	authToken.Type = "bearer"

	// create refresh token
	refreshToken := jwt.New(jwt.SigningMethodRS256)
	refreshTokenExpired := timeNow.Add(o.jwtTokenRefreshTimeExp * time.Minute).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims

	decodePrivateKeyRefresh, _ := base64.StdEncoding.DecodeString(config.Get().Auth.JwtToken.RefreshToken.PrivateKey)
	privateKeyRsaRefresh, err := rsaheader.ReadPrivateKeyFromEnv(decodePrivateKeyRefresh)
	if err != nil {
		log.Err(err).Msg("cannot read access private key token rsa from env")
		return tokendto.Token{}
	}

	tokenRefreshString, err := refreshToken.SignedString(privateKeyRsaRefresh)
	if err != nil {
		log.Err(err).Msg("err read access private key rsa")
		return tokendto.Token{}
	}

	authToken.RefreshToken = tokenRefreshString

	// save to db
	ctx := context.Background()
	tx, err := o.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	tokenentity := entities.TokenDetails{
		Token:     authToken.RefreshToken,
		UserID:    claims["id"].(string),
		ExpiresIn: refreshTokenExpired,
	}
	o.UserRepo.TokenDelete(ctx, tx, tokenentity.UserID)
	o.UserRepo.TokenCreate(ctx, tx, tokenentity)

	return tokendto.Token{
		Type:         authToken.Type,
		Token:        authToken.Token,
		RefreshToken: authToken.RefreshToken,
	}
}
