package dto

import (
	"toko-bangunan/src/modules/purchase-product/entities"
)

type PurchaseProductRes struct {
	ID             string `json:"id"`
	IdProduct      string `json:"id_product"`
	PurchaseAmount int32  `json:"purchase_amount"`
	PurchasePrice  int64  `json:"purchase_price"`
	PurchaseTotal  int64  `json:"purchase_total"`
	Status         string `json:"status"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

func EntitiesToResponse(productEntity entities.PurchaseProduct) PurchaseProductRes {
	return PurchaseProductRes{
		ID:             productEntity.ID,
		IdProduct:      productEntity.IdProduct,
		PurchaseAmount: productEntity.PurchaseAmount,
		PurchasePrice:  productEntity.PurchasePrice,
		PurchaseTotal:  productEntity.PurchaseTotal,
		Status:         productEntity.Status,
		CreatedAt:      productEntity.CreatedAt,
		UpdatedAt:      productEntity.UpdatedAt,
	}
}
