package controllers

import (
	"toko-bangunan/helpers"
	"toko-bangunan/models/dto"
	"toko-bangunan/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryProductController interface {
	FindALL(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type CategoryProductControllerImpl struct {
	CategoryProductService services.CategoryProductService
}

func NewCategoryProductControllerImpl(categoryProductService services.CategoryProductService) CategoryProductController {
	return &CategoryProductControllerImpl{CategoryProductService: categoryProductService}
}

func (controller *CategoryProductControllerImpl) FindALL(c *fiber.Ctx) error {
	categoryProducts := controller.CategoryProductService.FindALL(c.Context())
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "list data category product", categoryProducts))
}

func (controller *CategoryProductControllerImpl) FindById(c *fiber.Ctx) error {
	id := c.Params("id", "nil")
	categoryProduct := controller.CategoryProductService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "list data category product", categoryProduct))
}

func (controller *CategoryProductControllerImpl) Create(c *fiber.Ctx) error {
	categoryProduct := new(dto.CreateCategoryProductReq)

	if err := c.BodyParser(categoryProduct); err != nil {
		return err
	}

	categoryProductResult := controller.CategoryProductService.Create(c.Context(), *categoryProduct)

	return c.Status(fiber.StatusCreated).JSON(helpers.NewResponse("ok", fiber.StatusCreated, "created category product successfully", categoryProductResult))
}

func (controller *CategoryProductControllerImpl) Update(c *fiber.Ctx) error {
	categoryProduct := new(dto.UpdateCategoryProductReq)

	if err := c.BodyParser(categoryProduct); err != nil {
		return err
	}

	categoryProductResult := controller.CategoryProductService.Update(c.Context(), *categoryProduct, c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "updated category product successfully", categoryProductResult))
}

func (controller *CategoryProductControllerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("id", "null")
	controller.CategoryProductService.Delete(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "delete category product successfully", nil))
}
