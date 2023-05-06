package entities

type Product struct {
	ID                string `json:"id"`
	SupplierId        string `json:"supplier_id"`
	ProductCategoryId string `json:"category_product_id"`
	Name              string `json:"name"`
	SellingPrice      int64  `json:"selling_price"`
	StockProduct      int32  `json:"stock_product"`
	Image             string `json:"image"`
	CreatedAt         int64  `json:"created_at"`
	UpdatedAt         int64  `json:"updated_at"`
}
