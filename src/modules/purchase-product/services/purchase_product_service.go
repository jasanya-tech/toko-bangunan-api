package services

import (
	"context"
	"database/sql"
	"time"

	"toko-bangunan/infrastructures/db/transaction"
	"toko-bangunan/internal/protocols/http/exception"
	productrepo "toko-bangunan/src/modules/product/repositories"
	"toko-bangunan/src/modules/purchase-product/dto"
	"toko-bangunan/src/modules/purchase-product/entities"
	purchaseproductrepo "toko-bangunan/src/modules/purchase-product/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PurchaseProductService interface {
	FindALL(ctx context.Context) *[]dto.PurchaseProductRes
	FindById(ctx context.Context, id string) *dto.PurchaseProductRes
	Create(ctx context.Context, request dto.CreatePurchaseProductReq) *dto.PurchaseProductRes
	Update(ctx context.Context, request dto.UpdatePurchaseProductReq, id string) *dto.PurchaseProductRes
	Delete(ctx context.Context, id string)
}

type PurchaseProductServiceImpl struct {
	PurchaseProductRepo purchaseproductrepo.Repositories
	ProductRepository   productrepo.Repositories
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewPurchaseProductServiceImpl(
	purchaseProductRepo purchaseproductrepo.Repositories,
	productRepository productrepo.Repositories,
	db *sql.DB,
	validate *validator.Validate,
) PurchaseProductService {
	return &PurchaseProductServiceImpl{
		PurchaseProductRepo: purchaseProductRepo,
		ProductRepository:   productRepository,
		DB:                  db,
		Validate:            validate,
	}
}

func (service *PurchaseProductServiceImpl) FindALL(ctx context.Context) *[]dto.PurchaseProductRes {
	purchaseProducts := service.PurchaseProductRepo.FindAll(ctx, service.DB)
	var purchaseProductResponses []dto.PurchaseProductRes

	for _, purchaseProduct := range *purchaseProducts {
		purchaseProductResponse := dto.EntitiesToResponse(purchaseProduct)
		purchaseProductResponses = append(purchaseProductResponses, purchaseProductResponse)
	}

	return &purchaseProductResponses
}

func (service *PurchaseProductServiceImpl) FindById(ctx context.Context, id string) *dto.PurchaseProductRes {
	purchaseProduct, err := service.PurchaseProductRepo.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}

	purchaseProductResponse := dto.EntitiesToResponse(*purchaseProduct)

	return &purchaseProductResponse
}

func (service *PurchaseProductServiceImpl) Create(ctx context.Context, request dto.CreatePurchaseProductReq) *dto.PurchaseProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	_, err = service.ProductRepository.FindById(ctx, service.DB, request.IdProduct)
	if err != nil {
		panic(err)
	}

	purchaseProductEntity := &entities.PurchaseProduct{
		ID:             uuid.NewString(),
		IdProduct:      request.IdProduct,
		PurchaseAmount: request.PurchaseAmount,
		PurchasePrice:  request.PurchasePrice,
		PurchaseTotal:  request.PurchasePrice * int64(request.PurchaseAmount),
		Status:         "unpaid",
		CreatedAt:      time.Now().Unix(),
		UpdatedAt:      time.Now().Unix(),
	}

	purchaseProductEntityCreate, err := service.PurchaseProductRepo.Create(ctx, tx, purchaseProductEntity)
	if err != nil {
		panic(err)
	}

	purchaseProductResponse := dto.EntitiesToResponse(*purchaseProductEntityCreate)

	return &purchaseProductResponse
}

func (service *PurchaseProductServiceImpl) Update(ctx context.Context, request dto.UpdatePurchaseProductReq, id string) *dto.PurchaseProductRes {
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		panic(errValidate)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	purchaseProductFind, err := service.PurchaseProductRepo.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}

	if purchaseProductFind.Status == "paid" {
		panic(exception.BadRequestError{Message: "data cannot be changed because it has already made a payment"})
	}

	produtFind, err := service.ProductRepository.FindById(ctx, service.DB, purchaseProductFind.IdProduct)
	if err != nil {
		panic(err)
	}

	purchaseProductEntity := &entities.PurchaseProduct{
		ID:             purchaseProductFind.ID,
		IdProduct:      request.IdProduct,
		PurchaseAmount: request.PurchaseAmount,
		PurchasePrice:  request.PurchasePrice,
		PurchaseTotal:  request.PurchasePrice * int64(request.PurchaseAmount),
		Status:         request.Status,
		CreatedAt:      purchaseProductFind.CreatedAt,
		UpdatedAt:      time.Now().Unix(),
	}
	purchaseProductEntityUpdate, err := service.PurchaseProductRepo.Update(ctx, tx, purchaseProductEntity)
	if err != nil {
		panic(err)
	}

	if purchaseProductEntityUpdate.Status == "paid" {
		produtFind.StockProduct = produtFind.StockProduct + purchaseProductEntityUpdate.PurchaseAmount
		_, err = service.ProductRepository.Update(ctx, tx, produtFind)
		if err != nil {
			panic(err)
		}
	}

	purchaseProductResponse := dto.EntitiesToResponse(*purchaseProductEntityUpdate)

	return &purchaseProductResponse
}

func (service *PurchaseProductServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer transaction.Transaction(tx)

	productPurchaseFind, err := service.PurchaseProductRepo.FindById(ctx, service.DB, id)
	if err != nil {
		panic(err)
	}

	productFind, err := service.ProductRepository.FindById(ctx, service.DB, productPurchaseFind.IdProduct)
	if err != nil {
		panic(err)
	}

	if productPurchaseFind.Status == "paid" {
		productFind.StockProduct = productFind.StockProduct - productPurchaseFind.PurchaseAmount
		_, err := service.ProductRepository.Update(ctx, tx, productFind)
		if err != nil {
			panic(err)
		}
	}

	err = service.PurchaseProductRepo.Delete(ctx, tx, productPurchaseFind.ID)
	if err != nil {
		panic(err)
	}
}
