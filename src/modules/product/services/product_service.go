package services

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"toko-bangunan/infrastructures/db/transaction"
	"toko-bangunan/internal/utils/format"
	categoryproductdto "toko-bangunan/src/modules/category-product/dto"
	categoryproductentity "toko-bangunan/src/modules/category-product/entities"
	categoryproductrepo "toko-bangunan/src/modules/category-product/repositories"
	"toko-bangunan/src/modules/product/dto"
	productentity "toko-bangunan/src/modules/product/entities"
	productrepo "toko-bangunan/src/modules/product/repositories"
	supplierdto "toko-bangunan/src/modules/supplier/dto"
	supplierentity "toko-bangunan/src/modules/supplier/entities"
	supplierrepo "toko-bangunan/src/modules/supplier/repositories"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	FindALL(ctx context.Context) *[]dto.ProductRes
	FindById(ctx context.Context, id string) (*dto.ProductRes, error)
	Create(ctx context.Context, request dto.CreateProductReq) (*dto.ProductRes, error)
	Update(ctx context.Context, request dto.UpdateProductReq, id string) (*dto.ProductRes, error)
	Delete(ctx context.Context, id string) error
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
		productResponse := dto.ProductRes{
			ID:              product.ID,
			Name:            product.Name,
			SellingPrice:    product.SellingPrice,
			PurchasePrice:   product.PurchasePrice,
			StockProduct:    product.StockProduct,
			Image:           product.Image,
			CreatedAt:       product.CreatedAt,
			UpdatedAt:       product.UpdatedAt,
			Supplier:        supplierdto.SupplierRes(product.Supplier),
			ProductCategory: categoryproductdto.CategoryProductRes(product.ProductCategory),
		}
		productResponses = append(productResponses, productResponse)
	}
	return &productResponses
}

func (service *ProductServiceImpl) FindById(ctx context.Context, id string) (*dto.ProductRes, error) {
	product, err := service.ProductRepository.FindById(ctx, service.DB, id)
	if err != nil {
		return nil, err
	}
	productResponse := &dto.ProductRes{
		ID:              product.ID,
		Name:            product.Name,
		SellingPrice:    product.SellingPrice,
		PurchasePrice:   product.PurchasePrice,
		StockProduct:    product.StockProduct,
		Image:           product.Image,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
		Supplier:        supplierdto.SupplierRes(product.Supplier),
		ProductCategory: categoryproductdto.CategoryProductRes(product.ProductCategory),
	}

	return productResponse, nil
}

func (service *ProductServiceImpl) Create(ctx context.Context, request dto.CreateProductReq) (*dto.ProductRes, error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return nil, errValidate
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer transaction.Transaction(tx)

	findSupplier, err := service.SupplierRepository.FindById(ctx, service.DB, request.SupplierId)
	if err != nil {
		return nil, err
	}

	findCategoryProduct, err := service.CategoryProductRepository.FindById(ctx, service.DB, request.ProductCategoryId)
	if err != nil {
		return nil, err
	}

	productEntity := &productentity.Product{
		ID:              format.NewItemID("PRD", time.Now()).String(),
		Supplier:        supplierentity.Supplier(*findSupplier),
		ProductCategory: categoryproductentity.CategoryProduct(*findCategoryProduct),
		SellingPrice:    request.SellingPrice,
		PurchasePrice:   request.PurchasePrice,
		StockProduct:    request.StockProduct,
		Image:           request.Image,
		Name:            request.Name,
		CreatedAt:       time.Now().Unix(),
		UpdatedAt:       time.Now().Unix(),
	}

	productEntity, err = service.ProductRepository.Create(ctx, tx, productEntity)
	if err != nil {
		return nil, err
	}

	productRespon := &dto.ProductRes{
		ID:              productEntity.ID,
		Supplier:        supplierdto.SupplierRes(productEntity.Supplier),
		ProductCategory: categoryproductdto.CategoryProductRes(productEntity.ProductCategory),
		Name:            productEntity.Name,
		SellingPrice:    productEntity.SellingPrice,
		PurchasePrice:   productEntity.PurchasePrice,
		StockProduct:    productEntity.StockProduct,
		Image:           productEntity.Image,
		CreatedAt:       productEntity.CreatedAt,
		UpdatedAt:       productEntity.UpdatedAt,
	}
	return productRespon, nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, request dto.UpdateProductReq, id string) (*dto.ProductRes, error) {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return nil, errValidate
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer transaction.Transaction(tx)

	productFind, err := service.ProductRepository.FindById(ctx, service.DB, id)
	if err != nil {
		return nil, err
	}

	findSupplier, err := service.SupplierRepository.FindById(ctx, service.DB, request.SupplierId)
	if err != nil {
		return nil, err
	}

	findCategoryProduct, err := service.CategoryProductRepository.FindById(ctx, service.DB, request.ProductCategoryId)
	if err != nil {
		return nil, err
	}

	product := &productentity.Product{
		ID:            productFind.ID,
		Name:          request.Name,
		SellingPrice:  request.SellingPrice,
		PurchasePrice: request.PurchasePrice,
		StockProduct:  request.StockProduct,
		Image:         request.Image,
		CreatedAt:     productFind.CreatedAt,
		UpdatedAt:     time.Now().Unix(),
		Supplier: supplierentity.Supplier{
			ID:        request.SupplierId,
			Name:      findSupplier.Name,
			Email:     findSupplier.Email,
			Phone:     findSupplier.Phone,
			Address:   findSupplier.Address,
			CreatedAt: findSupplier.CreatedAt,
			UpdatedAt: findSupplier.UpdatedAt,
		},
		ProductCategory: categoryproductentity.CategoryProduct{
			ID:        request.ProductCategoryId,
			Name:      findCategoryProduct.Name,
			CreatedAt: findCategoryProduct.CreatedAt,
			UpdatedAt: findCategoryProduct.UpdatedAt,
		},
	}
	if request.Image != "" {
		os.Remove(fmt.Sprintf("./public/images/product-img/%s", productFind.Image))
	} else {
		product.Image = productFind.Image
	}

	product, err = service.ProductRepository.Update(ctx, tx, product)
	if err != nil {
		return nil, err
	}

	productResponse := &dto.ProductRes{
		ID:              product.ID,
		Name:            product.Name,
		SellingPrice:    product.SellingPrice,
		PurchasePrice:   product.PurchasePrice,
		StockProduct:    product.StockProduct,
		Image:           product.Image,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
		Supplier:        supplierdto.SupplierRes(product.Supplier),
		ProductCategory: categoryproductdto.CategoryProductRes(product.ProductCategory),
	}
	return productResponse, nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer transaction.Transaction(tx)
	productFind, err := service.ProductRepository.FindById(ctx, service.DB, id)
	if err != nil {
		return err
	}
	os.Remove(fmt.Sprintf("./public/images/product-img/%s", productFind.Image))

	service.ProductRepository.Delete(ctx, tx, productFind.ID)
	return nil
}
