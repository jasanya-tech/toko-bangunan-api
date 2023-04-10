package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/SyaibanAhmadRamadhan/toko-bangunan/helpers"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/models/dto"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/models/entities"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/repositories"
	"github.com/go-playground/validator/v10"
)

type CategoryProductService interface {
	Create(ctx context.Context, request dto.CreateCategoryProductReq) dto.CategoryProductRes
}

type CategoryProductServiceImpl struct {
	CategoryProductRepo repositories.CategoryProductRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewCategoryProductService(categortProductRepo repositories.CategoryProductRepository, db *sql.DB, validate *validator.Validate) CategoryProductService {
	return &CategoryProductServiceImpl{CategoryProductRepo: categortProductRepo, DB: db, Validate: validate}
}

func (service *CategoryProductServiceImpl) Create(ctx context.Context, request dto.CreateCategoryProductReq) dto.CategoryProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)

	categoryProduct := entities.CategoryProduct{
		ID:        request.ID,
		Name:      request.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	categoryProduct = service.CategoryProductRepo.Create(ctx, tx, categoryProduct)
	categoryProductRes := dto.CategoryProductRes{
		ID:        categoryProduct.ID,
		Name:      categoryProduct.Name,
		CreatedAt: categoryProduct.CreatedAt,
		UpdatedAt: categoryProduct.UpdatedAt,
	}
	return categoryProductRes
}
