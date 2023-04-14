package controllers

import (
	"toko-bangunan/helpers"
	"toko-bangunan/models/dto"
	"toko-bangunan/services"

	"github.com/gofiber/fiber/v2"
)

type SupplierController interface {
	FindALL(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type SupplierControllerImpl struct {
	SupplierService services.SupplierService
}

func NewSupplierControllerImpl(SupplierService services.SupplierService) SupplierController {
	return &SupplierControllerImpl{SupplierService: SupplierService}
}

func (controller *SupplierControllerImpl) FindALL(c *fiber.Ctx) error {
	suppliers := controller.SupplierService.FindALL(c.Context())
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "list data supplier", suppliers))
}

func (controller *SupplierControllerImpl) FindById(c *fiber.Ctx) error {
	supplier := controller.SupplierService.FindById(c.Context(), c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "data supplier", supplier))
}

func (controller *SupplierControllerImpl) Create(c *fiber.Ctx) error {
	supplier := new(dto.CreateSupplierReq)

	if err := c.BodyParser(supplier); err != nil {
		panic(err)
	}

	supplierSercice := controller.SupplierService.Create(c.Context(), *supplier)
	return c.Status(fiber.StatusCreated).JSON(helpers.NewResponse("ok", fiber.StatusCreated, "created supplier successfully", supplierSercice))
}

func (controller *SupplierControllerImpl) Update(c *fiber.Ctx) error {
	supplier := new(dto.UpdateSupplierReq)
	if err := c.BodyParser(supplier); err != nil {
		panic(err)
	}
	supplierSercice := controller.SupplierService.Update(c.Context(), *supplier, c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "update supplier successfully", supplierSercice))
}

func (controller *SupplierControllerImpl) Delete(c *fiber.Ctx) error {
	controller.SupplierService.Delete(c.Context(), c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "delete supplier successfully", nil))
}
