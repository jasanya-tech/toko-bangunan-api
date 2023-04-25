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
	FindById(ctx context.Context, id string) (*dto.CategoryProductRes, error)
	Create(ctx context.Context, request dto.CreateCategoryProductReq) (*dto.CategoryProductRes, error)
	Update(ctx context.Context, request dto.UpdateCategoryProductReq, id string) (*dto.CategoryProductRes, error)
	Delete(ctx context.Context, id string) error
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
		categoryProductResponse := dto.CategoryProductRes{
			ID:        categoryProduct.ID,
			Name:      categoryProduct.Name,
			CreatedAt: categoryProduct.CreatedAt,
			UpdatedAt: categoryProduct.UpdatedAt,
		}
		categoryProductResponses = append(categoryProductResponses, categoryProductResponse)
	}
	return &categoryProductResponses
}

func (service *CategoryProductServiceImpl) FindById(ctx context.Context, id string) (*dto.CategoryProductRes, error) {
	categoryProduct, err := service.CategoryProductRepo.FindById(ctx, service.DB, id)
	if err != nil {
		return nil, err
	}
	var categoryProductResponse dto.CategoryProductRes

	categoryProductResponse.ID = categoryProduct.ID
	categoryProductResponse.Name = categoryProduct.Name
	categoryProductResponse.CreatedAt = categoryProduct.CreatedAt
	categoryProductResponse.UpdatedAt = categoryProduct.UpdatedAt
	return &categoryProductResponse, nil
}

func (service *CategoryProductServiceImpl) Create(ctx context.Context, request dto.CreateCategoryProductReq) (*dto.CategoryProductRes, error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer transaction.Transaction(tx)

	categoryProduct := &entities.CategoryProduct{
		ID:        format.NewItemID("CTG-PRD", time.Now()).String(),
		Name:      request.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	categoryProduct, err = service.CategoryProductRepo.Create(ctx, tx, categoryProduct)
	if err != nil {
		return nil, err
	}

	categoryProductRes := dto.CategoryProductRes{
		ID:        categoryProduct.ID,
		Name:      categoryProduct.Name,
		CreatedAt: categoryProduct.CreatedAt,
		UpdatedAt: categoryProduct.UpdatedAt,
	}
	return &categoryProductRes, nil
}

func (service *CategoryProductServiceImpl) Update(ctx context.Context, request dto.UpdateCategoryProductReq, id string) (*dto.CategoryProductRes, error) {
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
		return nil, err
	}
	categoryProduct := &entities.CategoryProduct{
		ID:        categoryProductFind.ID,
		Name:      request.Name,
		CreatedAt: categoryProductFind.CreatedAt,
		UpdatedAt: time.Now().Unix(),
	}
	categoryProduct, err = service.CategoryProductRepo.Update(ctx, tx, categoryProduct)
	if err != nil {
		return nil, err
	}

	categoryProductRes := dto.CategoryProductRes{
		ID:        id,
		Name:      categoryProduct.Name,
		CreatedAt: categoryProduct.CreatedAt,
		UpdatedAt: categoryProduct.UpdatedAt,
	}
	return &categoryProductRes, nil
}

func (service *CategoryProductServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer transaction.Transaction(tx)

	if err := service.CategoryProductRepo.Delete(ctx, tx, id); err != nil {
		return err
	}
	return nil
}
