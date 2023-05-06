package services

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"toko-bangunan/infrastructures/db/transaction"
	"toko-bangunan/internal/utils/format"
	categoryproductrepo "toko-bangunan/src/modules/category-product/repositories"
	"toko-bangunan/src/modules/product/dto"
	productentity "toko-bangunan/src/modules/product/entities"
	productrepo "toko-bangunan/src/modules/product/repositories"
	supplierrepo "toko-bangunan/src/modules/supplier/repositories"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	FindALL(ctx context.Context) *[]dto.ProductRes
	FindById(ctx context.Context, id string) *dto.ProductRes
	Create(ctx context.Context, request dto.CreateProductReq) *dto.ProductRes
	Update(ctx context.Context, request dto.UpdateProductReq, id string) *dto.ProductRes
	Delete(ctx context.Context, id string)
}

type ProductServiceImpl struct {
	SupplierRepository        supplierrepo.Repositories
	CategoryProductRepository categoryproductrepo.Repositories
	ProductRepository         productrepo.Repositories
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewProductServiceImpl(
	supplierRepository supplierrepo.Repositories,
	categoryProductRepository categoryproductrepo.Repositories,
	productRepository productrepo.Repositories,
	db *sql.DB,
	validate *validator.Validate,
) ProductService {
	return &ProductServiceImpl{
		SupplierRepository:        supplierRepository,
		CategoryProductRepository: categoryProductRepository,
		ProductRepository:         productRepository,
		DB:                        db,
		Validate:                  validate,
	}
}

func (service *ProductServiceImpl) FindALL(ctx context.Context) *[]dto.ProductRes {
	products := service.ProductRepository.FindAll(ctx, service.DB)
	var productResponses []dto.ProductRes

	for _, product := range *products {
		productResponse := dto.EntitiesToResponse(product)
		productResponses = append(productResponses, productResponse)
	}

	return &productResponses
}

func (service *ProductServiceImpl) FindById(ctx context.Context, id string) *dto.ProductRes {
	product, err := service.ProductRepository.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}
	productResponse := dto.EntitiesToResponse(*product)

	return &productResponse
}

func (service *ProductServiceImpl) Create(ctx context.Context, request dto.CreateProductReq) *dto.ProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	findSupplier, err := service.SupplierRepository.FindById(ctx, service.DB, request.SupplierId)
	if err != nil {
		panic(err)
	}

	findCategoryProduct, err := service.CategoryProductRepository.FindById(ctx, service.DB, request.ProductCategoryId)
	if err != nil {
		panic(err)
	}

	productEntity := &productentity.Product{
		ID:                format.NewItemID("PRD", time.Now()).String(),
		SupplierId:        findSupplier.ID,
		ProductCategoryId: findCategoryProduct.ID,
		SellingPrice:      request.SellingPrice,
		Image:             request.Image,
		Name:              request.Name,
		CreatedAt:         time.Now().Unix(),
		UpdatedAt:         time.Now().Unix(),
	}

	productEntity, err = service.ProductRepository.Create(ctx, tx, productEntity)
	if err != nil {
		panic(err)
	}

	productResponse := dto.EntitiesToResponse(*productEntity)
	return &productResponse
}

func (service *ProductServiceImpl) Update(ctx context.Context, request dto.UpdateProductReq, id string) *dto.ProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	productFind, err := service.ProductRepository.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}

	_, err = service.SupplierRepository.FindById(ctx, service.DB, request.SupplierId)
	if err != nil {
		panic(err)
	}

	_, err = service.CategoryProductRepository.FindById(ctx, service.DB, request.ProductCategoryId)
	if err != nil {
		panic(err)
	}

	product := &productentity.Product{
		ID:                productFind.ID,
		SupplierId:        request.SupplierId,
		ProductCategoryId: request.ProductCategoryId,
		Name:              request.Name,
		SellingPrice:      request.SellingPrice,
		CreatedAt:         productFind.CreatedAt,
		UpdatedAt:         time.Now().Unix(),
	}
	if request.Image != "" {
		product.Image = request.Image
	} else {
		product.Image = productFind.Image
	}

	product, err = service.ProductRepository.Update(ctx, tx, product)
	if err != nil {
		panic(product)
	}

	if request.Image != "" {
		os.Remove(fmt.Sprintf("./public/images/product-img/%s", productFind.Image))
	}

	productResponse := dto.EntitiesToResponse(*product)
	return &productResponse
}

func (service *ProductServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)
	productFind, err := service.ProductRepository.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}
	os.Remove(fmt.Sprintf("./public/images/product-img/%s", productFind.Image))

	service.ProductRepository.Delete(ctx, tx, productFind.ID)
}
