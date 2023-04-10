package controllers

import (
	"net/http"

	"github.com/SyaibanAhmadRamadhan/toko-bangunan/models/dto"
	"github.com/SyaibanAhmadRamadhan/toko-bangunan/services"
	"github.com/gofiber/fiber/v2"
)

type CategoryProductController interface {
	Create(c *fiber.Ctx) error
}

type CategoryProductControllerImpl struct {
	CategoryProductService services.CategoryProductService
}

func NewCategoryProductControllerImpl(categoryProductService services.CategoryProductService) CategoryProductController {
	return &CategoryProductControllerImpl{CategoryProductService: categoryProductService}
}

func (controller *CategoryProductControllerImpl) Create(c *fiber.Ctx) error {
	categoryProduct := new(dto.CreateCategoryProductReq)

	if err := c.BodyParser(categoryProduct); err != nil {
		return err
	}

	categoryProductResult := controller.CategoryProductService.Create(c.Context(), *categoryProduct)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    http.StatusCreated,
		"message": "created category product successfully",
		"data":    categoryProductResult,
	})
}
