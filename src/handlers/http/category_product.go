package http

import (
	"toko-bangunan/internal/protocols/http/response"
	"toko-bangunan/src/modules/category-product/dto"
	

	"github.com/gofiber/fiber/v2"
)

func (http *HttpHandlerImpl) FindAllCategoryProduct(c *fiber.Ctx) error {
	categoryProducts := http.CategoryProductService.FindALL(c.Context())
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "list data category product", categoryProducts))
}

func (http *HttpHandlerImpl) FindByIdCategoryProduct(c *fiber.Ctx) error {
	id := c.Params("id", "nil")
	categoryProduct, err := http.CategoryProductService.FindById(c.Context(), id)
	if err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "list data category product", categoryProduct))
}

func (http *HttpHandlerImpl) CreateCategoryProduct(c *fiber.Ctx) error {
	categoryProduct := new(dto.CreateCategoryProductReq)

	if err := c.BodyParser(categoryProduct); err != nil {
		return err
	}

	categoryProductResult, err := http.CategoryProductService.Create(c.Context(), *categoryProduct)
	if err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusCreated).JSON(response.NewResponse("ok", fiber.StatusCreated, "created category product successfully", categoryProductResult))
}

func (http *HttpHandlerImpl) UpdateCategoryProduct(c *fiber.Ctx) error {
	categoryProduct := new(dto.UpdateCategoryProductReq)

	if err := c.BodyParser(categoryProduct); err != nil {
		return err
	}

	categoryProductResult, err := http.CategoryProductService.Update(c.Context(), *categoryProduct, c.Params("id"))
	if err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "updated category product successfully", categoryProductResult))
}

func (http *HttpHandlerImpl) DeleteCategoryProduct(c *fiber.Ctx) error {
	id := c.Params("id", "null")
	if err := http.CategoryProductService.Delete(c.Context(), id); err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "delete category product successfully", nil))
}
