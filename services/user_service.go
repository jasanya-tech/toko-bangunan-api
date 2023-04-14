package services

import (
	"context"
	"database/sql"
	"time"

	"toko-bangunan/helpers"
	"toko-bangunan/helpers/exception"
	"toko-bangunan/models/dto"
	"toko-bangunan/models/entities"
	"toko-bangunan/repositories"
	"toko-bangunan/utils"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	FindALL(ctx context.Context) []dto.SupplierRes
	FindById(ctx context.Context, id string) dto.SupplierRes
	Create(ctx context.Context, request dto.CreateUserReq) dto.UserRes
	Update(ctx context.Context, request dto.UpdateSupplierReq, id string) dto.SupplierRes
	Delete(ctx context.Context, id string)
	AuthLogin(ctx context.Context, request dto.LoginReq) dto.UserRes
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repositories.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: db, Validate: validate}
}

func (service *UserServiceImpl) FindALL(ctx context.Context) []dto.SupplierRes {
	// suppliers := service.SupplierRepository.FindAll(ctx, service.DB)
	// var suppliersResponse []dto.SupplierRes

	// for _, supplier := range suppliers {
	// 	supplierResponse := dto.SupplierRes{
	// 		ID:        supplier.ID,
	// 		Name:      supplier.Name,
	// 		Email:     supplier.Email,
	// 		Phone:     supplier.Phone,
	// 		Address:   supplier.Address,
	// 		CreatedAt: supplier.CreatedAt,
	// 		UpdatedAt: supplier.UpdatedAt,
	// 	}
	// 	suppliersResponse = append(suppliersResponse, supplierResponse)
	// }
	// return suppliersResponse
	return []dto.SupplierRes{}
}

func (service *UserServiceImpl) FindById(ctx context.Context, id string) dto.SupplierRes {
	// supplier := service.SupplierRepository.FindById(ctx, service.DB, id)
	// supplierResponse := dto.SupplierRes{
	// 	ID:        supplier.ID,
	// 	Name:      supplier.Name,
	// 	Email:     supplier.Email,
	// 	Phone:     supplier.Phone,
	// 	Address:   supplier.Address,
	// 	CreatedAt: supplier.CreatedAt,
	// 	UpdatedAt: supplier.UpdatedAt,
	// }
	// return supplierResponse
	return dto.SupplierRes{}
}

func (service *UserServiceImpl) Create(ctx context.Context, request dto.CreateUserReq) dto.UserRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)

	passwordHash, err := utils.HashPassword(request.Password)
	if err != nil {
		panic(err)
	}
	userEntity := entities.User{
		ID:        string(utils.NewItemID("USR", time.Now())),
		Username:  request.Username,
		Email:     request.Email,
		Role:      1,
		Image:     request.Image,
		Password:  passwordHash,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	userEntity = service.UserRepository.Create(ctx, tx, userEntity)
	userRespon := dto.UserRes{
		ID:        userEntity.ID,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Role:      userEntity.Role,
		Image:     userEntity.Image,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
	return userRespon
}

func (service *UserServiceImpl) Update(ctx context.Context, request dto.UpdateSupplierReq, id string) dto.SupplierRes {
	// errValidate := service.Validate.Struct(request)
	// if errValidate != nil {
	// 	panic(errValidate)
	// }
	// tx, err := service.DB.Begin()
	// if err != nil {
	// 	panic(errValidate)
	// }
	// defer helpers.Transaction(tx)

	// supplierFind := service.SupplierRepository.FindById(ctx, service.DB, id)

	// supplier := entities.Supplier{
	// 	ID:        supplierFind.ID,
	// 	Name:      request.Name,
	// 	Email:     request.Email,
	// 	Phone:     request.Phone,
	// 	Address:   request.Address,
	// 	CreatedAt: supplierFind.CreatedAt,
	// 	UpdatedAt: time.Now().Unix(),
	// }
	// supplier = service.SupplierRepository.Update(ctx, tx, supplier)

	// supplierResponse := dto.SupplierRes{
	// 	ID:        supplier.ID,
	// 	Name:      supplier.Name,
	// 	Email:     supplier.Email,
	// 	Phone:     supplier.Phone,
	// 	Address:   supplier.Address,
	// 	CreatedAt: supplier.CreatedAt,
	// 	UpdatedAt: supplier.UpdatedAt,
	// }
	return dto.SupplierRes{}
}

func (service *UserServiceImpl) Delete(ctx context.Context, id string) {
	// tx, err := service.DB.Begin()
	// if err != nil {
	// 	panic(err)
	// }
	// defer helpers.Transaction(tx)

	// service.SupplierRepository.Delete(ctx, tx, id)
}

func (service *UserServiceImpl) AuthLogin(ctx context.Context, request dto.LoginReq) dto.UserRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	user := service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	userResponse := dto.UserRes{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	checkPassword := utils.CheckPasswordHash(request.Password, user.Password)
	if !checkPassword {
		panic(exception.BadRequestError{Message: "invalid email or password"})
	}
	return userResponse
}
