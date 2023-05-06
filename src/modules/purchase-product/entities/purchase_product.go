package entities

type PurchaseProduct struct {
	ID             string `json:"id"`
	IdProduct      string `json:"id_product"`
	PurchaseAmount int32  `json:"purchase_amount"`
	PurchasePrice  int64  `json:"purchase_price"`
	PurchaseTotal  int64  `json:"purchase_total"`
	Status         string `json:"status"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}
