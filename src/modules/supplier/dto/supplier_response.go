package dto

import "toko-bangunan/src/modules/supplier/entities"

type SupplierRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func EntitiesToResponse(entity entities.Supplier) SupplierRes {
	return SupplierRes{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Phone:     entity.Phone,
		Address:   entity.Address,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
