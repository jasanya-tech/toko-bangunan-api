package dto

type CreateProductReq struct {
	SupplierId        string `validate:"required" json:"supplier_id" form:"supplier_id"`
	ProductCategoryId string `validate:"required" json:"category_product_id" form:"category_product_id"`
	Name              string `validate:"required,min=1,max=100" json:"name" form:"name"`
	SellingPrice      int64  `validate:"required,min=1,max=9223372036854775807,number" json:"selling_price" form:"selling_price"`
	PurchasePrice     int64  `validate:"required,min=1,max=9223372036854775807,number" json:"purchase_price" form:"purchase_price"`
	StockProduct      int32  `validate:"required,min=1,max=2147483647,number" json:"stock_product" form:"stock_product"`
	Image             string `json:"image" form:"image"`
}
type UpdateProductReq struct {
	SupplierId        string `validate:"required" json:"supplier_id" form:"supplier_id"`
	ProductCategoryId string `validate:"required" json:"category_product_id" form:"category_product_id"`
	Name              string `validate:"required,min=1,max=100" json:"name" form:"name"`
	SellingPrice      int64  `validate:"required,min=1,max=9223372036854775807,number" json:"selling_price" form:"selling_price"`
	PurchasePrice     int64  `validate:"required,min=1,max=9223372036854775807,number" json:"purchase_price" form:"purchase_price"`
	StockProduct      int32  `validate:"required,min=1,max=2147483647,number" json:"stock_product" form:"stock_product"`
	Image             string `json:"image" form:"image"`
}
