package dto

import (
	"toko-bangunan/src/modules/product/entities"
)

type ProductRes struct {
	ID                string `json:"id"`
	SupplierId        string `json:"supplier"`
	ProductCategoryId string `json:"category_product"`
	Name              string `json:"name"`
	SellingPrice      int64  `json:"selling_price"`
	StockProduct      int32  `json:"stock_product"`
	Image             string `json:"image"`
	CreatedAt         int64  `json:"created_at"`
	UpdatedAt         int64  `json:"updated_at"`
}

func EntitiesToResponse(productEntity entities.Product) ProductRes {
	return ProductRes{
		ID:                productEntity.ID,
		SupplierId:        productEntity.SupplierId,
		ProductCategoryId: productEntity.ProductCategoryId,
		Name:              productEntity.Name,
		SellingPrice:      productEntity.SellingPrice,
		StockProduct:      productEntity.StockProduct,
		Image:             productEntity.Image,
		CreatedAt:         productEntity.CreatedAt,
		UpdatedAt:         productEntity.UpdatedAt,
	}
}
