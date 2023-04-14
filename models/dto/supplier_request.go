package dto

type CreateSupplierReq struct {
	Name    string `validate:"required,min=1,max=100" json:"name"`
	Email   string `validate:"required,min=1,max=100,email" json:"email"`
	Phone   string `validate:"required,min=10,max=15" json:"phone"`
	Address string `validate:"required" json:"address"`
}
type UpdateSupplierReq struct {
	Name    string `validate:"required,min=1,max=100" json:"name"`
	Email   string `validate:"required,min=1,max=100,email" json:"email"`
	Phone   string `validate:"required,min=10,max=15" json:"phone"`
	Address string `validate:"required" json:"address"`
}
