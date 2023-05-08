package http

import (
	"toko-bangunan/internal/protocols/http/response"
	"toko-bangunan/src/modules/supplier/dto"

	"github.com/gofiber/fiber/v2"
)

func (http *HttpHandlerImpl) FindAllSupplier(c *fiber.Ctx) error {
	suppliers := http.SupplierService.FindALL(c.Context())
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "list data supplier", suppliers, nil))
}

func (http *HttpHandlerImpl) FindByIdSupplier(c *fiber.Ctx) error {
	supplier := http.SupplierService.FindById(c.Context(), c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "data supplier", supplier, nil))
}

func (http *HttpHandlerImpl) CreateSupplier(c *fiber.Ctx) error {
	supplier := new(dto.CreateSupplierReq)

	if err := c.BodyParser(supplier); err != nil {
		panic(err)
	}

	supplierSercice := http.SupplierService.Create(c.Context(), *supplier)

	return c.Status(fiber.StatusCreated).JSON(response.NewResponse("ok", fiber.StatusCreated, "created supplier successfully", supplierSercice, nil))
}

func (http *HttpHandlerImpl) UpdateSupplier(c *fiber.Ctx) error {
	supplier := new(dto.UpdateSupplierReq)
	if err := c.BodyParser(supplier); err != nil {
		panic(err)
	}
	supplierSercice := http.SupplierService.Update(c.Context(), *supplier, c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "update supplier successfully", supplierSercice, nil))
}

func (http *HttpHandlerImpl) DeleteSupplier(c *fiber.Ctx) error {
	http.SupplierService.Delete(c.Context(), c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "delete supplier successfully", nil, nil))
}
