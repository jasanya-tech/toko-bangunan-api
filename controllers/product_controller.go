package controllers

import (
	"fmt"

	"toko-bangunan/helpers"
	"toko-bangunan/helpers/exception"
	"toko-bangunan/models/dto"
	"toko-bangunan/services"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	FindALL(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductControllerImpl(productService services.ProductService) ProductController {
	return &ProductControllerImpl{ProductService: productService}
}

func (controller *ProductControllerImpl) FindALL(c *fiber.Ctx) error {
	products := controller.ProductService.FindALL(c.Context())
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "list data product", products))
}

func (controller *ProductControllerImpl) FindById(c *fiber.Ctx) error {
	product := controller.ProductService.FindById(c.Context(), c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "data product", product))
}

func (controller *ProductControllerImpl) Create(c *fiber.Ctx) error {
	product := new(dto.CreateProductReq)

	if err := c.BodyParser(product); err != nil {
		panic(err)
	}

	fileImg := c.Locals("filename")
	if fileImg == "" {
		panic(exception.BadRequestError{Message: "image required"})
	}
	product.Image = fileImg.(string)

	productService := controller.ProductService.Create(c.Context(), *product)

	file, _ := c.FormFile("image")
	if file != nil {
		if err := c.SaveFile(file, fmt.Sprintf("%s/%s", "./public/images/product-img", fileImg)); err != nil {
			panic(err)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(helpers.NewResponse("ok", fiber.StatusCreated, "created product successfully", productService))
}

func (controller *ProductControllerImpl) Update(c *fiber.Ctx) error {
	product := new(dto.UpdateProductReq)
	if err := c.BodyParser(product); err != nil {
		panic(err)
	}
	fileImg := c.Locals("filename")
	product.Image = fileImg.(string)

	productService := controller.ProductService.Update(c.Context(), *product, c.Params("id"))

	file, _ := c.FormFile("image")
	if file != nil {
		if err := c.SaveFile(file, fmt.Sprintf("%s/%s", "./public/images/product-img", fileImg)); err != nil {
			panic(err)
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "update product successfully", productService))
}

func (controller *ProductControllerImpl) Delete(c *fiber.Ctx) error {
	controller.ProductService.Delete(c.Context(), c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(helpers.NewResponse("ok", fiber.StatusOK, "delete product successfully", nil))
}
