package services

import (
	"context"
	"database/sql"
	"time"

	"toko-bangunan/helpers"
	"toko-bangunan/models/dto"
	"toko-bangunan/models/entities"
	"toko-bangunan/repositories"
	"toko-bangunan/utils"

	"github.com/go-playground/validator/v10"
)

type SupplierService interface {
	FindALL(ctx context.Context) []dto.SupplierRes
	FindById(ctx context.Context, id string) dto.SupplierRes
	Create(ctx context.Context, request dto.CreateSupplierReq) dto.SupplierRes
	Update(ctx context.Context, request dto.UpdateSupplierReq, id string) dto.SupplierRes
	Delete(ctx context.Context, id string)
}

type SupplierServiceImpl struct {
	SupplierRepository repositories.SupplierRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewSupplierServiceImpl(supplierRepository repositories.SupplierRepository, db *sql.DB, validate *validator.Validate) SupplierService {
	return &SupplierServiceImpl{SupplierRepository: supplierRepository, DB: db, Validate: validate}
}

func (service *SupplierServiceImpl) FindALL(ctx context.Context) []dto.SupplierRes {
	suppliers := service.SupplierRepository.FindAll(ctx, service.DB)
	var suppliersResponse []dto.SupplierRes

	for _, supplier := range suppliers {
		supplierResponse := dto.SupplierRes{
			ID:        supplier.ID,
			Name:      supplier.Name,
			Email:     supplier.Email,
			Phone:     supplier.Phone,
			Address:   supplier.Address,
			CreatedAt: supplier.CreatedAt,
			UpdatedAt: supplier.UpdatedAt,
		}
		suppliersResponse = append(suppliersResponse, supplierResponse)
	}
	return suppliersResponse
}

func (service *SupplierServiceImpl) FindById(ctx context.Context, id string) dto.SupplierRes {
	supplier := service.SupplierRepository.FindById(ctx, service.DB, id)
	supplierResponse := dto.SupplierRes{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Email:     supplier.Email,
		Phone:     supplier.Phone,
		Address:   supplier.Address,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
	return supplierResponse
}

func (service *SupplierServiceImpl) Create(ctx context.Context, request dto.CreateSupplierReq) dto.SupplierRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)

	supplierEntity := entities.Supplier{
		ID:        string(utils.NewItemID("SPR", time.Now())),
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Address:   request.Address,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	supplierEntity = service.SupplierRepository.Create(ctx, tx, supplierEntity)
	supplierRespon := dto.SupplierRes{
		ID:        supplierEntity.ID,
		Name:      supplierEntity.Name,
		Email:     supplierEntity.Email,
		Phone:     supplierEntity.Phone,
		Address:   supplierEntity.Address,
		CreatedAt: supplierEntity.CreatedAt,
		UpdatedAt: supplierEntity.UpdatedAt,
	}
	return supplierRespon
}

func (service *SupplierServiceImpl) Update(ctx context.Context, request dto.UpdateSupplierReq, id string) dto.SupplierRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(errValidate)
	}
	defer helpers.Transaction(tx)

	supplierFind := service.SupplierRepository.FindById(ctx, service.DB, id)

	supplier := entities.Supplier{
		ID:        supplierFind.ID,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Address:   request.Address,
		CreatedAt: supplierFind.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}
	supplier = service.SupplierRepository.Update(ctx, tx, supplier)

	supplierResponse := dto.SupplierRes{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Email:     supplier.Email,
		Phone:     supplier.Phone,
		Address:   supplier.Address,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
	return supplierResponse
}

func (service *SupplierServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)

	service.SupplierRepository.Delete(ctx, tx, id)
}
