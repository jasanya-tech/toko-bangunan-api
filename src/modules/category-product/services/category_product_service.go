package services

import (
	"context"
	"database/sql"
	"time"

	"toko-bangunan/infrastructures/db/transaction"
	"toko-bangunan/internal/utils/format"
	"toko-bangunan/src/modules/category-product/dto"
	"toko-bangunan/src/modules/category-product/entities"
	"toko-bangunan/src/modules/category-product/repositories"

	"github.com/go-playground/validator/v10"
)

type CategoryProductService interface {
	FindALL(ctx context.Context) *[]dto.CategoryProductRes
	FindById(ctx context.Context, id string) *dto.CategoryProductRes
	Create(ctx context.Context, request dto.CreateCategoryProductReq) *dto.CategoryProductRes
	Update(ctx context.Context, request dto.UpdateCategoryProductReq, id string) *dto.CategoryProductRes
	Delete(ctx context.Context, id string)
}

type CategoryProductServiceImpl struct {
	CategoryProductRepo repositories.Repositories
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewCategoryProductService(categortProductRepo repositories.Repositories, db *sql.DB, validate *validator.Validate) CategoryProductService {
	return &CategoryProductServiceImpl{CategoryProductRepo: categortProductRepo, DB: db, Validate: validate}
}

func (service *CategoryProductServiceImpl) FindALL(ctx context.Context) *[]dto.CategoryProductRes {
	categoryProducts := service.CategoryProductRepo.FindALL(ctx, service.DB)
	var categoryProductResponses []dto.CategoryProductRes

	for _, categoryProduct := range *categoryProducts {
		categoryProductResponse := dto.EntitiesToResponse(categoryProduct)
		categoryProductResponses = append(categoryProductResponses, categoryProductResponse)
	}
	return &categoryProductResponses
}

func (service *CategoryProductServiceImpl) FindById(ctx context.Context, id string) *dto.CategoryProductRes {
	categoryProduct, err := service.CategoryProductRepo.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}

	categoryProductResponse := dto.EntitiesToResponse(*categoryProduct)
	return &categoryProductResponse
}

func (service *CategoryProductServiceImpl) Create(ctx context.Context, request dto.CreateCategoryProductReq) *dto.CategoryProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	categoryProduct := &entities.CategoryProduct{
		ID:        format.NewItemID("CTG-PRD", time.Now()).String(),
		Name:      request.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	categoryProductCreate, err := service.CategoryProductRepo.Create(ctx, tx, categoryProduct)
	if err != nil {
		panic(err)
	}

	categoryProductResponse := dto.EntitiesToResponse(*categoryProductCreate)
	return &categoryProductResponse
}

func (service *CategoryProductServiceImpl) Update(ctx context.Context, request dto.UpdateCategoryProductReq, id string) *dto.CategoryProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	categoryProductFind, err := service.CategoryProductRepo.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}
	categoryProduct := &entities.CategoryProduct{
		ID:        categoryProductFind.ID,
		Name:      request.Name,
		CreatedAt: categoryProductFind.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}
	categoryProductUpdate, err := service.CategoryProductRepo.Update(ctx, tx, categoryProduct)
	if err != nil {
		panic(err)
	}

	categoryProductResponse := dto.EntitiesToResponse(*categoryProductUpdate)
	return &categoryProductResponse
}

func (service *CategoryProductServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	if err := service.CategoryProductRepo.Delete(ctx, tx, id); err != nil {
		panic(err)
	}
}
