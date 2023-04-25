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
	FindById(ctx context.Context, id string) (*dto.SupplierRes, error)
	Create(ctx context.Context, request dto.CreateSupplierReq) (*dto.SupplierRes, error)
	Update(ctx context.Context, request dto.UpdateSupplierReq, id string) (*dto.SupplierRes, error)
	Delete(ctx context.Context, id string) error
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
	return &suppliersResponse
}

func (service *SupplierServiceImpl) FindById(ctx context.Context, id string) (*dto.SupplierRes, error) {
	supplier, err := service.SupplierRepository.FindById(ctx, service.DB, id)
	if err != nil {
		return nil, err
	}
	supplierResponse := &dto.SupplierRes{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Email:     supplier.Email,
		Phone:     supplier.Phone,
		Address:   supplier.Address,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
	return supplierResponse, nil
}

func (service *SupplierServiceImpl) Create(ctx context.Context, request dto.CreateSupplierReq) (*dto.SupplierRes, error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return nil, errValidate
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
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
	supplierEntity, err = service.SupplierRepository.Create(ctx, tx, supplierEntity)
	if err != nil {
		return nil, err
	}
	supplierRespon := &dto.SupplierRes{
		ID:        supplierEntity.ID,
		Name:      supplierEntity.Name,
		Email:     supplierEntity.Email,
		Phone:     supplierEntity.Phone,
		Address:   supplierEntity.Address,
		CreatedAt: supplierEntity.CreatedAt,
		UpdatedAt: supplierEntity.UpdatedAt,
	}
	return supplierRespon, nil
}

func (service *SupplierServiceImpl) Update(ctx context.Context, request dto.UpdateSupplierReq, id string) (*dto.SupplierRes, error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return nil, errValidate
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer transaction.Transaction(tx)

	supplierFind, err := service.SupplierRepository.FindById(ctx, service.DB, id)
	if err != nil {
		return nil, err
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
	supplier, err = service.SupplierRepository.Update(ctx, tx, supplier)
	if err != nil {
		return nil, err
	}
	supplierResponse := &dto.SupplierRes{
		ID:        supplier.ID,
		Name:      supplier.Name,
		Email:     supplier.Email,
		Phone:     supplier.Phone,
		Address:   supplier.Address,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
	return supplierResponse, nil
}

func (service *SupplierServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer transaction.Transaction(tx)

	if err := service.SupplierRepository.Delete(ctx, tx, id); err != nil {
		return err
	}
	return nil
}
