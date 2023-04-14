package services

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"toko-bangunan/helpers"
	"toko-bangunan/models/dto"
	"toko-bangunan/models/entities"
	"toko-bangunan/repositories"
	"toko-bangunan/utils"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	FindALL(ctx context.Context) []dto.ProductRes
	FindById(ctx context.Context, id string) dto.ProductRes
	Create(ctx context.Context, request dto.CreateProductReq) dto.ProductRes
	Update(ctx context.Context, request dto.UpdateProductReq, id string) dto.ProductRes
	Delete(ctx context.Context, id string)
}

type ProductServiceImpl struct {
	SupplierRepository        repositories.SupplierRepository
	CategoryProductRepository repositories.CategoryProductRepository
	ProductRepository         repositories.ProductRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewProductServiceImpl(
	supplierRepository repositories.SupplierRepository,
	categoryProductRepository repositories.CategoryProductRepository,
	productRepository repositories.ProductRepository,
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

func (service *ProductServiceImpl) FindALL(ctx context.Context) []dto.ProductRes {
	products := service.ProductRepository.FindAll(ctx, service.DB)
	var productResponses []dto.ProductRes

	for _, product := range products {
		productResponse := dto.ProductRes{
			ID:              product.ID,
			Name:            product.Name,
			SellingPrice:    product.SellingPrice,
			PurchasePrice:   product.PurchasePrice,
			StockProduct:    product.StockProduct,
			Image:           product.Image,
			CreatedAt:       product.CreatedAt,
			UpdatedAt:       product.UpdatedAt,
			Supplier:        dto.SupplierRes(product.Supplier),
			ProductCategory: dto.CategoryProductRes(product.ProductCategory),
		}
		productResponses = append(productResponses, productResponse)
	}
	return productResponses
}

func (service *ProductServiceImpl) FindById(ctx context.Context, id string) dto.ProductRes {
	product := service.ProductRepository.FindById(ctx, service.DB, id)
	productResponse := dto.ProductRes{
		ID:              product.ID,
		Name:            product.Name,
		SellingPrice:    product.SellingPrice,
		PurchasePrice:   product.PurchasePrice,
		StockProduct:    product.StockProduct,
		Image:           product.Image,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
		Supplier:        dto.SupplierRes(product.Supplier),
		ProductCategory: dto.CategoryProductRes(product.ProductCategory),
	}

	return productResponse
}

func (service *ProductServiceImpl) Create(ctx context.Context, request dto.CreateProductReq) dto.ProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)
	findSupplier := service.SupplierRepository.FindById(ctx, service.DB, request.SupplierId)
	findCategoryProduct := service.CategoryProductRepository.FindById(ctx, service.DB, request.ProductCategoryId)

	productEntity := entities.Product{
		ID:              string(utils.NewItemID("PRD", time.Now())),
		Supplier:        entities.Supplier(findSupplier),
		ProductCategory: entities.CategoryProduct(findCategoryProduct),
		SellingPrice:    request.SellingPrice,
		PurchasePrice:   request.PurchasePrice,
		StockProduct:    request.StockProduct,
		Image:           request.Image,
		Name:            request.Name,
		CreatedAt:       time.Now().Unix(),
		UpdatedAt:       time.Now().Unix(),
	}
	productEntity = service.ProductRepository.Create(ctx, tx, productEntity)
	productRespon := dto.ProductRes{
		ID:              productEntity.ID,
		Supplier:        dto.SupplierRes(productEntity.Supplier),
		ProductCategory: dto.CategoryProductRes(productEntity.ProductCategory),
		Name:            productEntity.Name,
		SellingPrice:    productEntity.SellingPrice,
		PurchasePrice:   productEntity.PurchasePrice,
		StockProduct:    productEntity.StockProduct,
		Image:           productEntity.Image,
		CreatedAt:       productEntity.CreatedAt,
		UpdatedAt:       productEntity.UpdatedAt,
	}
	return productRespon
}

func (service *ProductServiceImpl) Update(ctx context.Context, request dto.UpdateProductReq, id string) dto.ProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(errValidate)
	}
	defer helpers.Transaction(tx)

	productFind := service.ProductRepository.FindById(ctx, service.DB, id)
	supplierFind := service.SupplierRepository.FindById(ctx, service.DB, request.SupplierId)
	categoryProductFind := service.CategoryProductRepository.FindById(ctx, service.DB, request.ProductCategoryId)

	product := entities.Product{
		ID:            productFind.ID,
		Name:          request.Name,
		SellingPrice:  request.SellingPrice,
		PurchasePrice: request.PurchasePrice,
		StockProduct:  request.StockProduct,
		Image:         request.Image,
		CreatedAt:     productFind.CreatedAt,
		UpdatedAt:     time.Now().Unix(),
		Supplier: entities.Supplier{
			ID:        request.SupplierId,
			Name:      supplierFind.Name,
			Email:     supplierFind.Email,
			Phone:     supplierFind.Phone,
			Address:   supplierFind.Address,
			CreatedAt: supplierFind.CreatedAt,
			UpdatedAt: supplierFind.UpdatedAt,
		},
		ProductCategory: entities.CategoryProduct{
			ID:        request.ProductCategoryId,
			Name:      categoryProductFind.Name,
			CreatedAt: categoryProductFind.CreatedAt,
			UpdatedAt: categoryProductFind.UpdatedAt,
		},
	}
	if request.Image != "" {
		os.Remove(fmt.Sprintf("./public/images/product-img/%s", productFind.Image))
	} else {
		product.Image = productFind.Image
	}
	product = service.ProductRepository.Update(ctx, tx, product)

	productResponse := dto.ProductRes{
		ID:              product.ID,
		Name:            product.Name,
		SellingPrice:    product.SellingPrice,
		PurchasePrice:   product.PurchasePrice,
		StockProduct:    product.StockProduct,
		Image:           product.Image,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
		Supplier:        dto.SupplierRes(product.Supplier),
		ProductCategory: dto.CategoryProductRes(product.ProductCategory),
	}
	return productResponse
}

func (service *ProductServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helpers.Transaction(tx)
	productFind := service.ProductRepository.FindById(ctx, service.DB, id)
	os.Remove(fmt.Sprintf("./public/images/product-img/%s", productFind.Image))

	service.ProductRepository.Delete(ctx, tx, id)
}
