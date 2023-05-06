package services

import (
	"context"
	"database/sql"
	"time"

	"toko-bangunan/infrastructures/db/transaction"
	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/internal/utils/auth"
	"toko-bangunan/internal/utils/hashing"
	userdto "toko-bangunan/src/modules/user/dto"
	userrepo "toko-bangunan/src/modules/user/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Create(ctx context.Context, request userdto.CreateUserReq) (*userdto.UserResponse, error)
	FindById(ctx context.Context, id string) (*userdto.UserResponse, error)
	FindTokenByUserId(ctx context.Context, userId string) (*userdto.TokenResponse, error)
	UserLogin(ctx context.Context, request userdto.LoginReq) (*userdto.TokenResponse, error)
	UserRefreshToken(ctx context.Context, userId string) (*userdto.TokenResponse, error)
	DeleteToken(ctx context.Context, userId string) error
}

type UserServiceImpl struct {
	UserRepository userrepo.Repositories
	DB             *sql.DB
	Validate       *validator.Validate
	JwtAuth        auth.JwtToken
}

func NewUserServiceImpl(
	userRepository userrepo.Repositories,
	db *sql.DB,
	validate *validator.Validate,
	jwtAuth auth.JwtToken,
) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
		JwtAuth:        jwtAuth,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request userdto.CreateUserReq) (response *userdto.UserResponse, error error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return nil, errValidate
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer transaction.Transaction(tx)

	passwordHash, err := hashing.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}
	request.Password = passwordHash

	userEntity := request.RequestToEntity(request)
	userCreate, err := service.UserRepository.Create(ctx, tx, userEntity)
	if err != nil {
		return nil, err
	}

	return response.EntityToResponse(userCreate), nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id string) (response *userdto.UserResponse, error error) {
	user, err := service.UserRepository.FindById(ctx, service.DB, id)
	if err != nil {
		return nil, err
	}

	return response.EntityToResponse(user), nil
}

func (service *UserServiceImpl) FindTokenByUserId(ctx context.Context, userId string) (*userdto.TokenResponse, error) {
	token, err := service.UserRepository.FindTokenByUserId(ctx, service.DB, userId)
	if err != nil {
		return nil, err
	}

	tokenResponse := userdto.TokenResponse{
		RefreshToken: token.Token,
	}
	return &tokenResponse, nil
}

func (service *UserServiceImpl) UserLogin(ctx context.Context, request userdto.LoginReq) (*userdto.TokenResponse, error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return nil, errValidate
	}

	user, err := service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	if err != nil {
		return nil, err
	}

	checkPassword := hashing.CheckPassword(request.Password, user.Password)
	if !checkPassword {
		return nil, exception.BadRequestError{Message: "invalid email or password"}
	}
	userToken := service.JwtAuth.SignRsa(jwt.MapClaims{
		"id": user.ID,
	})

	tokenRespon := userdto.TokenResponse{
		Type:         userToken.Type,
		UserId:       user.ID,
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}
	return &tokenRespon, nil
}

func (service *UserServiceImpl) UserRefreshToken(ctx context.Context, userId string) (*userdto.TokenResponse, error) {
	token, err := service.UserRepository.FindTokenByUserId(ctx, service.DB, userId)
	if err != nil {
		return nil, err
	}

	if token.ExpiresIn < time.Now().Unix() {
		return nil, exception.Unauthorize{Message: "refresh token expired"}
	}

	userToken := service.JwtAuth.SignRsa(jwt.MapClaims{
		"id": token.UserID,
	})

	tokenRespon := userdto.TokenResponse{
		Type:         userToken.Type,
		UserId:       token.UserID,
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}
	return &tokenRespon, nil
}

func (service *UserServiceImpl) DeleteToken(ctx context.Context, userId string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer transaction.Transaction(tx)
	service.UserRepository.TokenDelete(ctx, tx, userId)
	return nil
}
