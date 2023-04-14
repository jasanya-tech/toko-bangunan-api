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

type CategoryProductService interface {
	FindALL(ctx context.Context) []dto.CategoryProductRes
	FindById(ctx context.Context, id string) dto.CategoryProductRes
	Create(ctx context.Context, request dto.CreateCategoryProductReq) dto.CategoryProductRes
	Update(ctx context.Context, request dto.UpdateCategoryProductReq, id string) dto.CategoryProductRes
	Delete(ctx context.Context, id string)
}

type CategoryProductServiceImpl struct {
	CategoryProductRepo repositories.CategoryProductRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewCategoryProductService(categortProductRepo repositories.CategoryProductRepository, db *sql.DB, validate *validator.Validate) CategoryProductService {
	return &CategoryProductServiceImpl{CategoryProductRepo: categortProductRepo, DB: db, Validate: validate}
}

func (service *CategoryProductServiceImpl) FindALL(ctx context.Context) []dto.CategoryProductRes {
	categoryProducts := service.CategoryProductRepo.FindALL(ctx, service.DB)
	var categoryProductResponses []dto.CategoryProductRes

	for _, categoryProduct := range categoryProducts {
		categoryProductResponse := dto.CategoryProductRes{
			ID:        categoryProduct.ID,
			Name:      categoryProduct.Name,
			CreatedAt: categoryProduct.CreatedAt,
			UpdatedAt: categoryProduct.UpdatedAt,
		}
		categoryProductResponses = append(categoryProductResponses, categoryProductResponse)
	}
	return categoryProductResponses
}

func (service *CategoryProductServiceImpl) FindById(ctx context.Context, id string) dto.CategoryProductRes {
	categoryProduct := service.CategoryProductRepo.FindById(ctx, service.DB, id)
	var categoryProductResponse dto.CategoryProductRes

	categoryProductResponse.ID = categoryProduct.ID
	categoryProductResponse.Name = categoryProduct.Name
	categoryProductResponse.CreatedAt = categoryProduct.CreatedAt
	categoryProductResponse.UpdatedAt = categoryProduct.UpdatedAt
	return categoryProductResponse
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
		ID:        string(utils.NewItemID("CTG-PRD", time.Now())),
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

func (service *CategoryProductServiceImpl) Update(ctx context.Context, request dto.UpdateCategoryProductReq, id string) dto.CategoryProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)
	categoryProductFind := service.CategoryProductRepo.FindById(ctx, service.DB, id)
	categoryProduct := entities.CategoryProduct{
		ID:        categoryProductFind.ID,
		Name:      request.Name,
		CreatedAt: categoryProductFind.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}
	categoryProduct = service.CategoryProductRepo.Update(ctx, tx, categoryProduct)
	categoryProductRes := dto.CategoryProductRes{
		ID:        id,
		Name:      categoryProduct.Name,
		CreatedAt: categoryProduct.CreatedAt,
		UpdatedAt: categoryProduct.UpdatedAt,
	}
	return categoryProductRes
}

func (service *CategoryProductServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)

	service.CategoryProductRepo.Delete(ctx, tx, id)
}
