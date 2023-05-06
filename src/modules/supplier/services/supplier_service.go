package services

import (
	"context"
	"database/sql"
	"time"

	"toko-bangunan/infrastructures/db/transaction"
	"toko-bangunan/internal/utils/format"
	"toko-bangunan/src/modules/supplier/dto"
	"toko-bangunan/src/modules/supplier/entities"
	"toko-bangunan/src/modules/supplier/repositories"

	"github.com/go-playground/validator/v10"
)

type SupplierService interface {
	FindALL(ctx context.Context) *[]dto.SupplierRes
	FindById(ctx context.Context, id string) *dto.SupplierRes
	Create(ctx context.Context, request dto.CreateSupplierReq) *dto.SupplierRes
	Update(ctx context.Context, request dto.UpdateSupplierReq, id string) *dto.SupplierRes
	Delete(ctx context.Context, id string)
}

type SupplierServiceImpl struct {
	SupplierRepository repositories.Repositories
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewSupplierServiceImpl(supplierRepository repositories.Repositories, db *sql.DB, validate *validator.Validate) SupplierService {
	return &SupplierServiceImpl{SupplierRepository: supplierRepository, DB: db, Validate: validate}
}

func (service *SupplierServiceImpl) FindALL(ctx context.Context) *[]dto.SupplierRes {
	suppliers := service.SupplierRepository.FindAll(ctx, service.DB)
	var suppliersResponse []dto.SupplierRes

	for _, supplier := range *suppliers {
		supplierResponse := dto.EntitiesToResponse(supplier)
		suppliersResponse = append(suppliersResponse, supplierResponse)
	}
	return &suppliersResponse
}

func (service *SupplierServiceImpl) FindById(ctx context.Context, id string) *dto.SupplierRes {
	supplier, err := service.SupplierRepository.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}
	supplierResponse := dto.EntitiesToResponse(*supplier)
	return &supplierResponse
}

func (service *SupplierServiceImpl) Create(ctx context.Context, request dto.CreateSupplierReq) *dto.SupplierRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	supplierEntity := &entities.Supplier{
		ID:        format.NewItemID("SPR", time.Now()).String(),
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Address:   request.Address,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	supplierCreate, err := service.SupplierRepository.Create(ctx, tx, supplierEntity)
	if err != nil {
		panic(err)
	}

	supplierResponse := dto.EntitiesToResponse(*supplierCreate)
	return &supplierResponse
}

func (service *SupplierServiceImpl) Update(ctx context.Context, request dto.UpdateSupplierReq, id string) *dto.SupplierRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	supplierFind, err := service.SupplierRepository.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}

	supplier := &entities.Supplier{
		ID:        supplierFind.ID,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Address:   request.Address,
		CreatedAt: supplierFind.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}

	supplierUpdate, err := service.SupplierRepository.Update(ctx, tx, supplier)
	if err != nil {
		panic(err)
	}

	supplierResponse := dto.EntitiesToResponse(*supplierUpdate)
	return &supplierResponse
}

func (service *SupplierServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	if err := service.SupplierRepository.Delete(ctx, tx, id); err != nil {
		panic(err)
	}
}
