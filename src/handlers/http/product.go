package http

import (
	"fmt"

	"toko-bangunan/internal/protocols/http/exception"
	"toko-bangunan/internal/protocols/http/response"
	"toko-bangunan/src/modules/product/dto"

	"github.com/gofiber/fiber/v2"
)

func (http *HttpHandlerImpl) FindAllProduct(c *fiber.Ctx) error {
	products := http.ProductService.FindALL(c.Context())
	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "list data product", products))
}

func (http *HttpHandlerImpl) FindByIdProduct(c *fiber.Ctx) error {
	product := http.ProductService.FindById(c.Context(), c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "data product", product))
}

func (http *HttpHandlerImpl) CreateProduct(c *fiber.Ctx) error {
	product := new(dto.CreateProductReq)

	if err := c.BodyParser(product); err != nil {
		panic(err)
	}

	fileImg := c.Locals("filename")
	if fileImg == "" {
		panic(exception.BadRequestError{Message: "image required"})
	}
	product.Image = fileImg.(string)

	productService := http.ProductService.Create(c.Context(), *product)

	file, _ := c.FormFile("image")
	if file != nil {
		if err := c.SaveFile(file, fmt.Sprintf("%s/%s", "./public/images/product-img", fileImg)); err != nil {
			panic(err)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(response.NewResponse("ok", fiber.StatusCreated, "created product successfully", productService))
}

func (http *HttpHandlerImpl) UpdateProduct(c *fiber.Ctx) error {
	product := new(dto.UpdateProductReq)
	if err := c.BodyParser(product); err != nil {
		panic(err)
	}
	fileImg := c.Locals("filename")
	product.Image = fileImg.(string)

	productService := http.ProductService.Update(c.Context(), *product, c.Params("id"))

	file, _ := c.FormFile("image")
	if file != nil {
		if err := c.SaveFile(file, fmt.Sprintf("%s/%s", "./public/images/product-img", fileImg)); err != nil {
			panic(err)
		}
	}

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "update product successfully", productService))
}

func (http *HttpHandlerImpl) DeleteProduct(c *fiber.Ctx) error {
	http.ProductService.Delete(c.Context(), c.Params("id"))

	return c.Status(fiber.StatusOK).JSON(response.NewResponse("ok", fiber.StatusOK, "delete product successfully", nil))
}
