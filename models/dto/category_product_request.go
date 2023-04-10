package dto

type CreateCategoryProductReq struct {
	ID   string `json:"id"`
	Name string `validate:"required" json:"name"`
}
