package dto

import (
	categoryproductdto "toko-bangunan/src/modules/category-product/dto"
	supplierdto "toko-bangunan/src/modules/supplier/dto"
)

type ProductRes struct {
	ID              string                                `json:"id"`
	Name            string                                `json:"name"`
	SellingPrice    int64                                 `json:"selling_price"`
	PurchasePrice   int64                                 `json:"purchase_price"`
	StockProduct    int32                                 `json:"stock_product"`
	Image           string                                `json:"image"`
	CreatedAt       int64                                 `json:"created_at"`
	UpdatedAt       int64                                 `json:"updated_at"`
	Supplier        supplierdto.SupplierRes               `json:"supplier"`
	ProductCategory categoryproductdto.CategoryProductRes `json:"category_product"`
}
