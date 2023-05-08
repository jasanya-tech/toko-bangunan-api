package http

import (
	"toko-bangunan/internal/protocols/http/response"
	"toko-bangunan/src/modules/purchase-product/dto"

	"github.com/gofiber/fiber/v2"
)

func (http *HttpHandlerImpl) FindAllPurchaseProduct(c *fiber.Ctx) error {
	purchaseProducts := http.PurchaseProductService.FindALL(c.Context())
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "list data product", purchaseProducts, nil))
}

func (http *HttpHandlerImpl) FindByIdPurchaseProduct(c *fiber.Ctx) error {
	purchaseProduct := http.PurchaseProductService.FindById(c.Context(), c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "data product", purchaseProduct, nil))
}

func (http *HttpHandlerImpl) CreatePurchaseProduct(c *fiber.Ctx) error {
	purchaseProduct := new(dto.CreatePurchaseProductReq)

	if err := c.BodyParser(purchaseProduct); err != nil {
		panic(err)
	}

	purchaseProductService := http.PurchaseProductService.Create(c.Context(), *purchaseProduct)

	return c.Status(fiber.StatusCreated).JSON(response.NewResponse("ok", fiber.StatusCreated, "created purchase product successfully", purchaseProductService, nil))
}

func (http *HttpHandlerImpl) UpdatePurchaseProduct(c *fiber.Ctx) error {
	productPurchase := new(dto.UpdatePurchaseProductReq)
	if err := c.BodyParser(productPurchase); err != nil {
		panic(err)
	}
	productPuchaseService := http.PurchaseProductService.Update(c.Context(), *productPurchase, c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "update product successfully", productPuchaseService, nil))
}

func (http *HttpHandlerImpl) DeletePurchaseProduct(c *fiber.Ctx) error {
	http.PurchaseProductService.Delete(c.Context(), c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "delete purchase product successfully", nil, nil))
}
