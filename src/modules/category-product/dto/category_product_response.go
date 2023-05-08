package dto

import (
	categoryproductentity "toko-bangunan/src/modules/category-product/entities"
)

type CategoryProductRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func EntitiesToResponse(entity categoryproductentity.CategoryProduct) CategoryProductRes {
	return CategoryProductRes{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
