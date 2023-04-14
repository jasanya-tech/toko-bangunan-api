package dto

type CreateUserReq struct {
	Username string `validate:"required,min=1,max=255" json:"username"`
	Email    string `validate:"required,min=1,max=100,email" json:"email"`
	Image    string `json:"image"`
	Password string `validate:"required" json:"password"`
}

type LoginReq struct {
	Email    string `validate:"required,min=1,max=100,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
