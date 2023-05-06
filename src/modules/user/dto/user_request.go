package dto

import (
	"time"

	"toko-bangunan/internal/utils/format"
	userentity "toko-bangunan/src/modules/user/entities"
)

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

func (create CreateUserReq) RequestToEntity(request CreateUserReq) *userentity.User {
	return &userentity.User{
		ID:        format.NewItemID("USR", time.Now()).String(),
		Username:  request.Username,
		Email:     request.Email,
		Role:      1,
		Image:     request.Image,
		Password:  request.Password,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
