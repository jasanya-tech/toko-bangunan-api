package dto

import (
	userentity "toko-bangunan/src/modules/user/entities"
)

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      int8   `json:"role"`
	Image     string `json:"image"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type TokenResponse struct {
	Type         string `json:"type"`
	UserId       string `json:"id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (user *UserResponse) EntityToResponse(entitiy *userentity.User) *UserResponse {
	return &UserResponse{
		ID:        entitiy.ID,
		Username:  entitiy.Username,
		Email:     entitiy.Email,
		Role:      entitiy.Role,
		Image:     entitiy.Image,
		CreatedAt: entitiy.CreatedAt,
		UpdatedAt: entitiy.UpdatedAt,
	}
}
