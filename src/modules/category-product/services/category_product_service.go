package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"toko-bangunan/infrastructures/db/transaction"
	"toko-bangunan/internal/helpers/pagination"
	paginateentity "toko-bangunan/internal/helpers/pagination/entities"
	"toko-bangunan/internal/utils/format"
	"toko-bangunan/src/modules/category-product/dto"
	"toko-bangunan/src/modules/category-product/entities"
	"toko-bangunan/src/modules/category-product/repositories"

	"github.com/go-playground/validator/v10"
)

type CategoryProductService interface {
	FindALL(ctx context.Context, paginate paginateentity.Pagination) ([]dto.CategoryProductRes, paginateentity.PaginationInfo)
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

func (service *CategoryProductServiceImpl) FindALL(ctx context.Context, paginate paginateentity.Pagination) ([]dto.CategoryProductRes, paginateentity.PaginationInfo) {
	var whereStr string
	isFirstPage := paginate.Cursor == ""
	pointNext := false

	if paginate.Cursor != "" {
		decodedCursor, _ := pagination.DecodedCursor(paginate.Cursor)
		pointNext = decodedCursor["point_next"] == true
		operation, order := pagination.GetPaginationOperator(pointNext, paginate.Order)
		createdAt := int64(decodedCursor["created_at"].(float64))
		if order != "" {
			paginate.Order = order
		}
		// panic(createdAt)
		whereStr = fmt.Sprintf("WHERE created_at %s %d ORDER BY created_at %s LIMIT %d", operation, createdAt, paginate.Order, paginate.Page+1)
	} else {
		whereStr = fmt.Sprintf("ORDER BY created_at %s LIMIT %d", paginate.Order, paginate.Page+1)
	}

	categoryProducts := service.CategoryProductRepo.FindALL(ctx, service.DB, whereStr)
	var categoryProductResponses []dto.CategoryProductRes
	var paginateCalculateEntities []paginateentity.PaginationCalculate
	for _, categoryProduct := range *categoryProducts {
		categoryProductResponse := dto.EntitiesToResponse(categoryProduct)
		categoryProductResponses = append(categoryProductResponses, categoryProductResponse)

		paginateCalculateEntities = append(paginateCalculateEntities, paginateentity.PaginationCalculate{
			ID: categoryProduct.ID, CreatedAt: categoryProduct.CreatedAt,
		})
	}

	hasPagination := len(categoryProductResponses) > paginate.Page
	if hasPagination {
		categoryProductResponses = categoryProductResponses[:paginate.Page]
	}

	if !isFirstPage && !pointNext {
		categoryProductResponses = pagination.Reverse(categoryProductResponses)
	}

	pageInfo := pagination.CalculatePagination(isFirstPage, hasPagination, paginate.Page, paginateCalculateEntities, pointNext)
	return categoryProductResponses, pageInfo
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
