package http

import (
	"strconv"

	"toko-bangunan/internal/protocols/http/response"
	"toko-bangunan/internal/helpers/pagination/entities"
	"toko-bangunan/src/modules/category-product/dto"

	"github.com/gofiber/fiber/v2"
)

func (http *HttpHandlerImpl) FindAllCategoryProduct(c *fiber.Ctx) error {
	perPage := c.Query("per_page", "10")
	sortOrder := c.Query("sort_order", "asc")
	cursor := c.Query("cursor", "")
	limit, _ := strconv.Atoi(perPage)
	if limit < 1 || limit > 100 {
		limit = 10
	}

	paginateEntity := entities.Pagination{
		Page:   limit,
		Order:  sortOrder,
		Cursor: cursor,
	}

	categoryProduct, paginate := http.CategoryProductService.FindALL(c.Context(), paginateEntity)

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "list data category product", categoryProduct, &paginate))
}

func (http *HttpHandlerImpl) FindByIdCategoryProduct(c *fiber.Ctx) error {
	id := c.Params("id", "nil")
	categoryProduct := http.CategoryProductService.FindById(c.Context(), id)

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "list data category product", categoryProduct, nil))
}

func (http *HttpHandlerImpl) CreateCategoryProduct(c *fiber.Ctx) error {
	categoryProduct := new(dto.CreateCategoryProductReq)

	if err := c.BodyParser(categoryProduct); err != nil {
		return err
	}

	categoryProductResult := http.CategoryProductService.Create(c.Context(), *categoryProduct)

	return c.Status(fiber.StatusCreated).JSON(response.NewResponse("ok", fiber.StatusCreated, "created category product successfully", categoryProductResult, nil))
}

func (http *HttpHandlerImpl) UpdateCategoryProduct(c *fiber.Ctx) error {
	categoryProduct := new(dto.UpdateCategoryProductReq)

	if err := c.BodyParser(categoryProduct); err != nil {
		return err
	}

	categoryProductResult := http.CategoryProductService.Update(c.Context(), *categoryProduct, c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "updated category product successfully", categoryProductResult, nil))
}

func (http *HttpHandlerImpl) DeleteCategoryProduct(c *fiber.Ctx) error {
	id := c.Params("id", "null")
	http.CategoryProductService.Delete(c.Context(), id)

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "delete category product successfully", nil, nil))
}
