package dto

type CreatePurchaseProductReq struct {
	IdProduct      string `validate:"required" json:"id_product" form:"id_product"`
	PurchasePrice  int64  `validate:"required,min=1,max=9223372036854775807,number" json:"purchase_price" form:"purchase_price"`
	PurchaseAmount int32  `validate:"required,min=1,max=2147483647,number" json:"purchase_amount" form:"purchase_amount"`
}

type UpdatePurchaseProductReq struct {
	IdProduct      string `validate:"required" json:"id_product" form:"id_product"`
	PurchasePrice  int64  `validate:"required,min=1,max=9223372036854775807,number" json:"purchase_price" form:"purchase_price"`
	PurchaseAmount int32  `validate:"required,min=1,max=2147483647,number" json:"purchase_amount" form:"purchase_amount"`
	Status         string `json:"status" form:"status"`
}
