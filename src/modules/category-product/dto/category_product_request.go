package dto

type CreateCategoryProductReq struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}
type UpdateCategoryProductReq struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}
